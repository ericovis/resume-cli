package resume

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	_ "embed"

	"gopkg.in/yaml.v2"
)

//go:embed templates/resume.yml
var template []byte

type Resume struct {
	Basics struct {
		Name     string `yaml:"name"`
		Label    string `yaml:"label"`
		Image    string `yaml:"image"`
		Email    string `yaml:"email"`
		Phone    string `yaml:"phone"`
		URL      string `yaml:"url"`
		Summary  string `yaml:"summary"`
		Location struct {
			Address     string `yaml:"address"`
			PostalCode  string `yaml:"postalCode"`
			City        string `yaml:"city"`
			CountryCode string `yaml:"countryCode"`
			Region      string `yaml:"region"`
		} `yaml:"location"`
		Profiles []struct {
			Network  string `yaml:"network"`
			Username string `yaml:"username"`
			URL      string `yaml:"url"`
		} `yaml:"profiles"`
	} `yaml:"basics"`
	Work []struct {
		Name       string   `yaml:"name"`
		Position   string   `yaml:"position"`
		URL        string   `yaml:"url"`
		StartDate  string   `yaml:"startDate"`
		EndDate    string   `yaml:"endDate"`
		Summary    string   `yaml:"summary"`
		Highlights []string `yaml:"highlights"`
	} `yaml:"work"`
	Volunteer []struct {
		Organization string   `yaml:"organization"`
		Position     string   `yaml:"position"`
		URL          string   `yaml:"url"`
		StartDate    string   `yaml:"startDate"`
		EndDate      string   `yaml:"endDate"`
		Summary      string   `yaml:"summary"`
		Highlights   []string `yaml:"highlights"`
	} `yaml:"volunteer"`
	Education []struct {
		Institution string   `yaml:"institution"`
		URL         string   `yaml:"url"`
		Area        string   `yaml:"area"`
		StudyType   string   `yaml:"studyType"`
		StartDate   string   `yaml:"startDate"`
		EndDate     string   `yaml:"endDate"`
		Score       string   `yaml:"score"`
		Courses     []string `yaml:"courses"`
	} `yaml:"education"`
	Awards []struct {
		Title   string `yaml:"title"`
		Date    string `yaml:"date"`
		Awarder string `yaml:"awarder"`
		Summary string `yaml:"summary"`
	} `yaml:"awards"`
	Certificates []struct {
		Name   string `yaml:"name"`
		Date   string `yaml:"date"`
		Issuer string `yaml:"issuer"`
		URL    string `yaml:"url"`
	} `yaml:"certificates"`
	Publications []struct {
		Name        string `yaml:"name"`
		Publisher   string `yaml:"publisher"`
		ReleaseDate string `yaml:"releaseDate"`
		URL         string `yaml:"url"`
		Summary     string `yaml:"summary"`
	} `yaml:"publications"`
	Skills []struct {
		Name     string   `yaml:"name"`
		Level    string   `yaml:"level"`
		Keywords []string `yaml:"keywords"`
	} `yaml:"skills"`
	Languages []struct {
		Language string `yaml:"language"`
		Fluency  string `yaml:"fluency"`
	} `yaml:"languages"`
	Interests []struct {
		Name     string   `yaml:"name"`
		Keywords []string `yaml:"keywords"`
	} `yaml:"interests"`
	References []struct {
		Name      string `yaml:"name"`
		Reference string `yaml:"reference"`
	} `yaml:"references"`
	Projects []struct {
		Name       string   `yaml:"name"`
		StartDate  string   `yaml:"startDate"`
		EndDate    string   `yaml:"endDate"`
		Summary    string   `yaml:"summary"`
		Highlights []string `yaml:"highlights"`
		URL        string   `yaml:"url"`
	} `yaml:"projects"`
}

func (r *Resume) fromJSON(path string) error {
	file, _ := os.ReadFile(path)

	err := json.Unmarshal([]byte(file), r)
	if err != nil {
		return err
	}

	return nil
}

func (r *Resume) fromYAML(path string) error {
	file, _ := os.ReadFile(path)

	err := yaml.Unmarshal([]byte(file), r)
	if err != nil {
		return err
	}

	return nil
}

func (r *Resume) Load(path string) error {
	if _, err := os.Stat(path); err != nil {
		return fmt.Errorf("file %s does not exist", path)
	}

	ext := filepath.Ext(path)

	if ext == ".json" {
		return r.fromJSON(path)
	} else if ext == ".yaml" || ext == ".yml" {
		return r.fromYAML(path)
	}

	return fmt.Errorf("invalid file type: %s", ext)
}

func FindResumeFileOnDir(dir string) (string, error) {
	var file string
	resumeDataFiles := []string{"resume.json", "resume.yaml", "resume.yml"}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return file, err
	}

	for _, f := range resumeDataFiles {
		if _, err := os.Stat(filepath.Join(dir, f)); err == nil {
			file = filepath.Join(dir, f)
			break
		}
	}

	if file == "" {
		return file, fmt.Errorf("no resume data file found on path: %s", dir)
	}

	return file, nil
}

func (r *Resume) toJSON() ([]byte, error) {
	return json.MarshalIndent(r, "", " ")
}

func (r *Resume) toYAML() ([]byte, error) {
	return yaml.Marshal(r)
}

func (r *Resume) Save(path string) error {
	var data []byte
	var err error
	ext := filepath.Ext(path)

	if ext == ".json" {
		data, err = r.toJSON()
	} else if ext == ".yaml" || ext == ".yml" {
		data, err = r.toYAML()
	} else {
		return fmt.Errorf("invalid file type: %s", ext)
	}

	if err == nil {
		fmt.Println("Saving resume to", path)
		err = os.WriteFile(path, data, 0755)
	}

	return err
}

func (r *Resume) Init(dir string, json bool) error {
	var path string
	err := yaml.Unmarshal(template, r)

	if err != nil {
		return err
	}

	if json {
		path = filepath.Join(dir, "resume.json")
	} else {
		path = filepath.Join(dir, "resume.yml")
	}

	return r.Save(path)
}
