package features

import "fmt"

type Feature struct {
	Name        string
	Description string
	Enabled     bool
}

type FeatureSet struct {
	Features     map[string]Feature
	everythingOn bool
}

func NewFeatureSet() *FeatureSet {
	return &FeatureSet{
		Features:     make(map[string]Feature),
		everythingOn: false,
	}
}

func (fs *FeatureSet) AddFeature(name string, description string, enabled bool) {
	fs.Features[name] = Feature{
		Name:        name,
		Description: description,
		Enabled:     enabled,
	}
}

func (fs *FeatureSet) IsEnabled(name string) bool {
	// If everythingOn is true, all features are enabled
	if fs.everythingOn {
		return true
	}

	feature, exists := fs.Features[name]
	if !exists {
		return false
	}
	return feature.Enabled
}

func (fs *FeatureSet) EnableFeatures(names []string) error {
	for _, name := range names {
		err := fs.EnableFeature(name)
		if err != nil {
			return err
		}
	}
	return nil
}

func (fs *FeatureSet) EnableFeature(name string) error {
	// Special case for "everything"
	if name == "everything" {
		fs.everythingOn = true
		return nil
	}

	feature, exists := fs.Features[name]
	if !exists {
		return fmt.Errorf("feature %s does not exist", name)
	}
	feature.Enabled = true
	fs.Features[name] = feature
	return nil
}
