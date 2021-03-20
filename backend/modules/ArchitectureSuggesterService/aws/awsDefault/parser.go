package awsDefault

import (
	awsdefaulPb "alfred/modules/ArchitectureSuggesterService/aws/awsDefault/pb"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"path/filepath"
)

func ParserData(file string) (*awsdefaulPb.AwsDefaultData, error) {
	filename, _ := filepath.Abs("./modules/ArchitectureSuggesterService/aws/awsDefault/" + file)

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var awsDefaultData awsdefaulPb.AwsDefaultData

	err = yaml.Unmarshal(yamlFile, &awsDefaultData)
	if err != nil {
		return nil, err
	}
	return &awsDefaultData, nil
}
