package models_test

import (
	"encoding/json"

	"github.com/cloudfoundry-incubator/diego-ssh/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Diego SSH Route", func() {
	var route models.SSHRoute

	BeforeEach(func() {
		route = models.SSHRoute{
			ContainerPort:   2222,
			HostFingerprint: "my-key-fingerprint",
			User:            "user",
			Password:        "password",
			PrivateKey:      "PEM_ENCODED_KEY",
		}
	})

	Describe("JSON Marshalling", func() {
		Context("when the user and password are missing", func() {
			var expectedJson string

			BeforeEach(func() {
				route.User = ""
				route.Password = ""

				expectedJson = `{
					"container_port": 2222,
					"host_fingerprint": "my-key-fingerprint",
					"private_key": "PEM_ENCODED_KEY"
				}`
			})

			It("marshals the structure correctly", func() {
				payload, err := json.Marshal(route)
				Ω(err).ShouldNot(HaveOccurred())

				Ω(payload).Should(MatchJSON(expectedJson))
			})
		})

		Context("when the private key and host fingerprint are empty", func() {
			var expectedJson string

			BeforeEach(func() {
				route.PrivateKey = ""
				route.HostFingerprint = ""

				expectedJson = `{
					"container_port": 2222,
					"user": "user",
					"password": "password"
				}`
			})

			It("marshals the structure correctly", func() {
				payload, err := json.Marshal(route)
				Ω(err).ShouldNot(HaveOccurred())

				Ω(payload).Should(MatchJSON(expectedJson))
			})
		})
	})

	Describe("Round Trip Marshalling", func() {
		It("successfully marshals and unmarshals", func() {
			payload, err := json.Marshal(route)
			Ω(err).ShouldNot(HaveOccurred())

			var result models.SSHRoute
			err = json.Unmarshal(payload, &result)
			Ω(err).ShouldNot(HaveOccurred())

			Ω(result).Should(Equal(route))
		})
	})
})
