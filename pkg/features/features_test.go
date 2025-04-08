package features

import (
	"testing"
)

func TestNewFeatureSet(t *testing.T) {
	fs := NewFeatureSet()
	if fs == nil {
		t.Fatal("Expected NewFeatureSet to return a non-nil pointer")
	}
	if fs.Features == nil {
		t.Fatal("Expected Features map to be initialized")
	}
	if len(fs.Features) != 0 {
		t.Fatalf("Expected Features map to be empty, got %d items", len(fs.Features))
	}
	if fs.everythingOn {
		t.Fatal("Expected everythingOn to be initialized as false")
	}
}

func TestAddFeature(t *testing.T) {
	fs := NewFeatureSet()

	// Test adding a feature
	fs.AddFeature("test-feature", "A test feature", true)

	// Verify feature was added correctly
	if len(fs.Features) != 1 {
		t.Errorf("Expected 1 feature, got %d", len(fs.Features))
	}

	feature, exists := fs.Features["test-feature"]
	if !exists {
		t.Fatal("Feature was not added to the map")
	}

	if feature.Name != "test-feature" {
		t.Errorf("Expected feature name to be 'test-feature', got '%s'", feature.Name)
	}

	if feature.Description != "A test feature" {
		t.Errorf("Expected feature description to be 'A test feature', got '%s'", feature.Description)
	}

	if !feature.Enabled {
		t.Error("Expected feature to be enabled")
	}

	// Test adding another feature
	fs.AddFeature("another-feature", "Another test feature", false)

	if len(fs.Features) != 2 {
		t.Errorf("Expected 2 features, got %d", len(fs.Features))
	}

	// Test overriding existing feature
	fs.AddFeature("test-feature", "Updated description", false)

	feature = fs.Features["test-feature"]
	if feature.Description != "Updated description" {
		t.Errorf("Expected feature description to be updated to 'Updated description', got '%s'", feature.Description)
	}

	if feature.Enabled {
		t.Error("Expected feature to be disabled after update")
	}
}

func TestIsEnabled(t *testing.T) {
	fs := NewFeatureSet()

	// Test with non-existent feature
	if fs.IsEnabled("non-existent") {
		t.Error("Expected IsEnabled to return false for non-existent feature")
	}

	// Test with disabled feature
	fs.AddFeature("disabled-feature", "A disabled feature", false)
	if fs.IsEnabled("disabled-feature") {
		t.Error("Expected IsEnabled to return false for disabled feature")
	}

	// Test with enabled feature
	fs.AddFeature("enabled-feature", "An enabled feature", true)
	if !fs.IsEnabled("enabled-feature") {
		t.Error("Expected IsEnabled to return true for enabled feature")
	}
}

func TestEnableFeature(t *testing.T) {
	fs := NewFeatureSet()

	// Test enabling non-existent feature
	err := fs.EnableFeature("non-existent")
	if err == nil {
		t.Error("Expected error when enabling non-existent feature")
	}

	// Test enabling feature
	fs.AddFeature("test-feature", "A test feature", false)

	if fs.IsEnabled("test-feature") {
		t.Error("Expected feature to be disabled initially")
	}

	err = fs.EnableFeature("test-feature")
	if err != nil {
		t.Errorf("Expected no error when enabling feature, got: %v", err)
	}

	if !fs.IsEnabled("test-feature") {
		t.Error("Expected feature to be enabled after EnableFeature call")
	}

	// Test enabling already enabled feature
	err = fs.EnableFeature("test-feature")
	if err != nil {
		t.Errorf("Expected no error when enabling already enabled feature, got: %v", err)
	}
}

func TestEnableFeatures(t *testing.T) {
	fs := NewFeatureSet()

	// Prepare features
	fs.AddFeature("feature1", "Feature 1", false)
	fs.AddFeature("feature2", "Feature 2", false)

	// Test enabling multiple features
	err := fs.EnableFeatures([]string{"feature1", "feature2"})
	if err != nil {
		t.Errorf("Expected no error when enabling features, got: %v", err)
	}

	if !fs.IsEnabled("feature1") {
		t.Error("Expected feature1 to be enabled")
	}

	if !fs.IsEnabled("feature2") {
		t.Error("Expected feature2 to be enabled")
	}

	// Test with non-existent feature in the list
	err = fs.EnableFeatures([]string{"feature1", "non-existent"})
	if err == nil {
		t.Error("Expected error when enabling list with non-existent feature")
	}

	// Test with empty list
	err = fs.EnableFeatures([]string{})
	if err != nil {
		t.Errorf("Expected no error with empty feature list, got: %v", err)
	}

	// Test enabling everything through EnableFeatures
	fs = NewFeatureSet()
	err = fs.EnableFeatures([]string{"everything"})
	if err != nil {
		t.Errorf("Expected no error when enabling 'everything', got: %v", err)
	}

	if !fs.everythingOn {
		t.Error("Expected everythingOn to be true after enabling 'everything' via EnableFeatures")
	}
}

func TestEnableEverything(t *testing.T) {
	fs := NewFeatureSet()

	// Add a disabled feature
	fs.AddFeature("test-feature", "A test feature", false)

	// Verify it's disabled
	if fs.IsEnabled("test-feature") {
		t.Error("Expected feature to be disabled initially")
	}

	// Enable "everything"
	err := fs.EnableFeature("everything")
	if err != nil {
		t.Errorf("Expected no error when enabling 'everything', got: %v", err)
	}

	// Verify everythingOn was set
	if !fs.everythingOn {
		t.Error("Expected everythingOn to be true after enabling 'everything'")
	}

	// Verify the previously disabled feature is now enabled
	if !fs.IsEnabled("test-feature") {
		t.Error("Expected feature to be enabled when everythingOn is true")
	}

	// Verify a non-existent feature is also enabled
	if !fs.IsEnabled("non-existent") {
		t.Error("Expected non-existent feature to be enabled when everythingOn is true")
	}
}

func TestIsEnabledWithEverythingOn(t *testing.T) {
	fs := NewFeatureSet()

	// Enable "everything"
	err := fs.EnableFeature("everything")
	if err != nil {
		t.Errorf("Expected no error when enabling 'everything', got: %v", err)
	}

	// Test that any feature name returns true with IsEnabled
	if !fs.IsEnabled("some-feature") {
		t.Error("Expected IsEnabled to return true for any feature when everythingOn is true")
	}

	if !fs.IsEnabled("another-feature") {
		t.Error("Expected IsEnabled to return true for any feature when everythingOn is true")
	}
}
