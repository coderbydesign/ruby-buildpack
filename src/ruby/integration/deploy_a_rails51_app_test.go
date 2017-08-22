package integration_test

import (
	"path/filepath"
	"time"

	"github.com/cloudfoundry/libbuildpack/cutlass"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rails 5.1 (Webpack/Yarn) App", func() {
	var app *cutlass.App

	AfterEach(func() {
		if app != nil {
			app.Destroy()
		}
		app = nil
	})

	BeforeEach(func() {
		app = cutlass.New(filepath.Join(bpDir, "cf_spec", "fixtures", "rails51"))
	})

	It("Installs node6 and runs", func() {
		PushAppAndConfirm(app)
		Expect(app.Stdout.String()).To(ContainSubstring("Installing node 6."))

		Expect(app.GetBody("/")).To(ContainSubstring("Hello World"))
		Eventually(func() string { return app.Stdout.String() }, 10*time.Second).Should(ContainSubstring(`Started GET "/" for`))
	})
})