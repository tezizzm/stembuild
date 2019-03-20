package commandparser_test

import (
	"context"
	"errors"
	"flag"
	. "github.com/cloudfoundry-incubator/stembuild/commandparser"
	"github.com/cloudfoundry-incubator/stembuild/commandparser/commandparserfakes"
	"github.com/google/subcommands"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("construct", func() {
	// Focus of this test is not to test the Flags.Parse functionality as much
	// as to test that the command line flags values are stored in the expected
	// struct variables. This adds a bit of protection when renaming flag parameters.
	Describe("SetFlags", func() {

		var f *flag.FlagSet
		var ConstrCmd *ConstructCmd

		BeforeEach(func() {
			f = flag.NewFlagSet("test", flag.ExitOnError)
			ConstrCmd = &ConstructCmd{}
			ConstrCmd.SetFlags(f)
		})

		var longformArgs = []string{
			"-winrm-ip", "10.0.0.5",
			"-winrm-username", "Admin",
			"-winrm-password", "some_password",
		}

		It("stores the value of a WinRM user", func() {
			err := f.Parse(longformArgs)
			Expect(err).ToNot(HaveOccurred())
			Expect(ConstrCmd.GetWinRMUser()).To(Equal("Admin"))
		})

		It("stores the value of a WinRM password ", func() {
			err := f.Parse(longformArgs)
			Expect(err).ToNot(HaveOccurred())
			Expect(ConstrCmd.GetWinRMPwd()).To(Equal("some_password"))
		})

		It("stores the value of the a WinRM IP", func() {
			err := f.Parse(longformArgs)
			Expect(err).ToNot(HaveOccurred())
			Expect(ConstrCmd.GetWinRMIp()).To(Equal("10.0.0.5"))
		})
	})

	Describe("Execute", func() {

		var f *flag.FlagSet
		var gf *GlobalFlags
		var ConstrCmd ConstructCmd
		var emptyContext context.Context

		var fakeFactory *commandparserfakes.FakeVMPreparerFactory
		var fakeVmConstruct *commandparserfakes.FakeVmConstruct
		var fakeValidator *commandparserfakes.FakeConstructCmdValidator
		var fakeMessenger *commandparserfakes.FakeConstructMessenger

		BeforeEach(func() {
			f = flag.NewFlagSet("test", flag.ExitOnError)
			gf = &GlobalFlags{false, false, false}

			fakeFactory = &commandparserfakes.FakeVMPreparerFactory{}
			fakeVmConstruct = &commandparserfakes.FakeVmConstruct{}
			fakeValidator = &commandparserfakes.FakeConstructCmdValidator{}
			fakeMessenger = &commandparserfakes.FakeConstructMessenger{}
			fakeFactory.VMPreparerReturns(fakeVmConstruct)

			ConstrCmd.SetFlags(f)
			ConstrCmd = NewConstructCmd(fakeFactory, fakeValidator, fakeMessenger)
			ConstrCmd.GlobalFlags = gf
			emptyContext = context.Background()
		})

		It("should execute the construct VM command", func() {
			fakeValidator.PopulatedArgsReturns(true)
			fakeValidator.LGPOInDirectoryReturns(true)
			fakeVmConstruct.CanConnectToVMReturns(nil)

			exitStatus := ConstrCmd.Execute(emptyContext, f)

			Expect(exitStatus).To(Equal(subcommands.ExitSuccess))
			Expect(fakeValidator.PopulatedArgsCallCount()).To(Equal(1))
			Expect(fakeValidator.LGPOInDirectoryCallCount()).To(Equal(1))

			Expect(fakeVmConstruct.CanConnectToVMCallCount()).To(Equal(1))
			Expect(fakeVmConstruct.PrepareVMCallCount()).To(Equal(1))

		})

		Context("with missing arguments", func() {
			It("should return an error", func() {
				fakeValidator.PopulatedArgsReturns(false)

				exitStatus := ConstrCmd.Execute(emptyContext, f)

				Expect(exitStatus).To(Equal(subcommands.ExitFailure))
				Expect(fakeMessenger.ArgumentsNotProvidedCallCount()).To(Equal(1))
			})
		})

		Context("with LGPO.zip not in current directory", func() {
			It("should return an error", func() {
				fakeValidator.PopulatedArgsReturns(true)
				fakeValidator.LGPOInDirectoryReturns(false)

				exitStatus := ConstrCmd.Execute(emptyContext, f)

				Expect(exitStatus).To(Equal(subcommands.ExitFailure))
				Expect(fakeMessenger.LGPONotFoundCallCount()).To(Equal(1))
			})
		})

		Context("with VM being unreachable", func() {
			It("should return an error", func() {
				fakeValidator.PopulatedArgsReturns(true)
				fakeValidator.LGPOInDirectoryReturns(true)

				connectionError := errors.New("can't reach the vm")
				fakeVmConstruct.CanConnectToVMReturns(connectionError)

				exitStatus := ConstrCmd.Execute(emptyContext, f)

				Expect(exitStatus).To(Equal(subcommands.ExitFailure))
				arg := fakeMessenger.CannotConnectToVMArgsForCall(0)
				Expect(arg).To(Equal(connectionError))
				Expect(fakeMessenger.CannotConnectToVMCallCount()).To(Equal(1))
			})
		})

		Context("with an error during VMPrepare", func() {
			It("should retrun an error", func() {
				fakeValidator.PopulatedArgsReturns(true)
				fakeValidator.LGPOInDirectoryReturns(true)
				fakeVmConstruct.PrepareVMReturns(errors.New("some error"))

				exitStatus := ConstrCmd.Execute(emptyContext, f)

				Expect(exitStatus).To(Equal(subcommands.ExitFailure))
				Expect(fakeMessenger.CannotPrepareVMCallCount()).To(Equal(1))
			})
		})
	})
})
