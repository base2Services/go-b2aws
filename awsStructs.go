package b2aws

import (
		"time"
	)

type RegionInfo struct {
	RegionName string `xml:"regionName"`
	RegionEndpoint string `xml:"regionEndpoint"`
}

type AwsRequest struct {
	RequestId string `xml:"requestId"`
}

type Regions struct {
	AwsRequest
	RegionInfo []RegionInfo `xml:"regionInfo>item"`
}

type Tag struct {
	Key string `xml:"key"`
	Value string `xml:"value"`
}

type Instance struct { // just add what ever you need here when you need it.
	InstanceId string `xml:"instanceId",json:"Id"`
	Architecture string `xml:"architecture"`
	InstanceType string `xml:"instanceType"`
	PrivateDnsName string `xml:"privateDnsName"`
	Tags []Tag `xml:"tagSet>item"`
	IpAddress string `xml:"ipAddress"`
	InstanceState string `xml:"instanceState>name"`
	CurrentInstanceState string `xml:"currentState>name"`
	PreviousInstanceState string `xml:"previousState>name"`
	InstanceStateCode string `xml:"instanceState>code"`
	ProfileName string // custom field
	Region string // custom field
}

type Group struct {
	GroupId string `xml:"groupId"`
	GroupName string `xml:"groupName"`
}

type Reservation struct {
	ReservationId string `xml:"reservationId"`
	OwnerId string `xml:"ownerId"`
	Groups []Group `xml:"groupSet>item"`
	Instances []Instance `xml:"instancesSet>item"`
}

type InstantStatuses struct {
	AwsRequest
	Instances []Instance `xml:"instanceStatusSet>item"`
}

type Instances struct {
	AwsRequest
	Reservations []Reservation `xml:"reservationSet>item"`
}

type StartInstance struct {
	AwsRequest
	Instances []Instance `xml:"instancesSet>item"`
}

type DescribeDBSnapshotsResponse struct {
	RequestId string `xml:"ResponseMetadata>RequestId"`
	DescribeDBSnapshotResult []DBSnapshot `xml:"DescribeDBSnapshotsResult>DBSnapshots>DBSnapshot"`
}

type DBSnapshot struct {
	Port int
	OptionGroupName string
	Engine string
	Status string
	SnapshotType string
	LicenseModel string
	EngineVersion string
	DBInstanceIdentifier string
	DBSnapshotIdentifier string
	SnapshotCreateTime time.Time
	AvailabilityZone string
	InstanceCreateTime time.Time
	PercentProgress int
	AllocatedStorage int
	MasterUsername string
}

