package b2aws

import (
		"net/url"
		"net/http"
		awsauth "github.com/smartystreets/go-aws-auth"
		"io/ioutil"
		"encoding/xml"
		"strconv"
		"log"
)

func GetRegions(accessKey string, secretKey string, client *http.Client) (regionMap map[string]string, err error) {
	regionMap = make(map[string]string)

	var treq *http.Request
	treq, err = http.NewRequest("GET", "https://ec2.ap-southeast-2.amazonaws.com/?Action=DescribeRegions&Version=2014-06-15", nil)
	if err != nil  { return regionMap, err }
	awscred := awsauth.Credentials{
		AccessKeyID: accessKey,
		SecretAccessKey: secretKey,
	}
	awsauth.Sign(treq, awscred)
	var tresp *http.Response
	tresp, err = client.Do(treq)
	if err != nil  { return regionMap, err }
	var regions Regions = Regions{}
	body,_ := ioutil.ReadAll(tresp.Body)
	if err = xml.Unmarshal(body, &regions); err != nil  { return }

	for _, each := range regions.RegionInfo {
		regionMap[each.RegionName] = each.RegionEndpoint
	}

	return 
}

func GetS3Object(accessKey string, secretKey string, bucketObjectUrl string, client *http.Client) (body []byte, err error) {
	treq, err := http.NewRequest("GET", bucketObjectUrl, nil)
	if err != nil { return }
	awscred := awsauth.Credentials{
		AccessKeyID: accessKey,
		SecretAccessKey: secretKey,
	}
	awsauth.Sign(treq, awscred)
	tresp, err := client.Do(treq)
	if err != nil  { return }

	body,err = ioutil.ReadAll(tresp.Body)

	return 
}

func GetInstances(accessKey string, secretKey string, regionEndpoint string, client *http.Client) (instanceList []Instance, err error) {
	instanceList = []Instance{}
	var treq *http.Request
	treq, err = http.NewRequest("GET", "https://"+regionEndpoint+"/?Action=DescribeInstances&Version=2014-06-15", nil)
	if err != nil { return }
	awscred := awsauth.Credentials{
		AccessKeyID: accessKey,
		SecretAccessKey: secretKey,
	}
	awsauth.Sign(treq, awscred)
	var tresp *http.Response
	tresp, err = client.Do(treq)
	if err != nil  { return }
	var instances Instances = Instances{}
	body,_ := ioutil.ReadAll(tresp.Body)
	if err = xml.Unmarshal(body, &instances); err != nil  { return }

	for _, reservation := range instances.Reservations {
		instanceList = append(instanceList, reservation.Instances...)
	}

	return 
}

func MultiInstancesURL(regionEndpoint string, action string, instanceIds ...string) (ourl string, err error) {
	params := url.Values{}
	params,err = url.ParseQuery("Action=" + action + "&Version=2014-06-15")
	if err != nil { return }
	for index,id := range instanceIds {
		params.Add("InstanceId." + strconv.Itoa(index),id)
	}
	ourl ="https://" + regionEndpoint + "/?" + params.Encode() 
	return
}

func StartInstances(accessKey string, secretKey string, regionEndpoint string, client *http.Client, instanceIds ...string) (instances StartInstance, err error, rurl string) {
	var treq *http.Request
	rurl, err = MultiInstancesURL(regionEndpoint, "StartInstances", instanceIds...)
	treq, err = http.NewRequest("GET", rurl, nil)
	if err != nil { return }
	awscred := awsauth.Credentials{
		AccessKeyID: accessKey,
		SecretAccessKey: secretKey,
	}
	awsauth.Sign(treq, awscred)
	var tresp *http.Response
	tresp, err = client.Do(treq)
	if err != nil  { return }
	instances = StartInstance{}
	body,_ := ioutil.ReadAll(tresp.Body)
	if err = xml.Unmarshal(body, &instances); err != nil  { return }

	return 
}

func StopInstances(accessKey string, secretKey string, regionEndpoint string, client *http.Client, instanceIds ...string) (instances StartInstance, err error, rurl string) {
	var treq *http.Request
	rurl, err = MultiInstancesURL(regionEndpoint, "StopInstances", instanceIds...)
	treq, err = http.NewRequest("GET", rurl, nil)
	if err != nil { return }
	awscred := awsauth.Credentials{
		AccessKeyID: accessKey,
		SecretAccessKey: secretKey,
	}
	awsauth.Sign(treq, awscred)
	var tresp *http.Response
	tresp, err = client.Do(treq)
	if err != nil  { return }
	instances = StartInstance{}
	body,_ := ioutil.ReadAll(tresp.Body)
	if err = xml.Unmarshal(body, &instances); err != nil  { return }

	return 
}

func GetInstancesStatus(accessKey string, secretKey string, regionEndpoint string, client *http.Client, all bool, instanceIds ...string) (instantStatuses InstantStatuses, err error) {
	var treq *http.Request
	rurl, err := MultiInstancesURL(regionEndpoint, "DescribeInstanceStatus", instanceIds...)
	if all {
		rurl = rurl+"&IncludeAllInstances=true"
	}
	treq, err = http.NewRequest("GET", rurl, nil)
	awscred := awsauth.Credentials{
		AccessKeyID: accessKey,
		SecretAccessKey: secretKey,
	}
	awsauth.Sign(treq, awscred)
	var tresp *http.Response
	tresp, err = client.Do(treq)
	if err != nil  { return }
	instantStatuses = InstantStatuses{}
	body,_ := ioutil.ReadAll(tresp.Body)
	if err = xml.Unmarshal(body, &instantStatuses); err != nil  { return }
	return
}

func GetRDSSnapshots(accessKey string, secretKey string, regionEndpoint string, client *http.Client, params string) (dbSnapshots DescribeDBSnapshotsResponse, err error) {
	var treq *http.Request
	treq, err = http.NewRequest("GET", "https://"+regionEndpoint+"/?Action=DescribeDBSnapshots&Version=2014-09-01", nil)
	if err != nil  { return dbSnapshots, err }
	awscred := awsauth.Credentials{
		AccessKeyID: accessKey,
		SecretAccessKey: secretKey,
	}
	awsauth.Sign(treq, awscred)
	var tresp *http.Response
	tresp, err = client.Do(treq)
	if err != nil  { return }
	dbSnapshots = DescribeDBSnapshotsResponse {}
	body,_ := ioutil.ReadAll(tresp.Body)
	if err = xml.Unmarshal(body, &dbSnapshots); err != nil  { return }
	return
}

