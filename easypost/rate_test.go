package easypost

import "testing"

func TestServices(t *testing.T) {
	t.Run("returns services for carrier", func(t *testing.T) {
		rates := Rates{
			{Carrier: "c1", Service: "s1"},
			{Carrier: "c2", Service: "s2"},
			{Carrier: "c1", Service: "s3"},
			{Carrier: "c1", Service: "s4"},
			{Carrier: "c2", Service: "s5"},
			{Carrier: "c1", Service: "s6"},
			{Carrier: "c1", Service: "s7"},
		}

		services := rates.Services("c2")

		for _, service := range services {
			if service != "s2" && service != "s5" {
				t.Fatal("unexpected service in list", service)
			}
		}
	})
}

func TestFilter(t *testing.T) {
	t.Run("returns all elements that satisfy the predicate", func(t *testing.T) {
		rates := Rates{
			{Carrier: "c1"},
			{Carrier: "c2"},
			{Carrier: "c3"},
			{Carrier: "c1"},
			{Carrier: "c1"},
		}

		filtered := rates.Filter(func(r *Rate) bool { return r.Carrier == "c1" })
		if len(filtered) != 3 {
			t.Fatal("expected length of 3 got", len(filtered))
		}

		for _, rate := range filtered {
			if rate.Carrier != "c1" {
				t.Fatalf("unexpected element %#v", rate)
			}
		}
	})
}
