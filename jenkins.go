package jenkins

import (
	"context"
	"github.com/bndr/gojenkins"
)

var ctx = context.Background()

func NewClient(address, username, password string) (*Client, error) {
	jenkins, err := gojenkins.CreateJenkins(nil, address, username, password).Init(ctx)
	if err != nil {
		return nil, err
	}
	return &Client{
		Address:  address,
		Username: username,
		Password: password,
		Jenkins:  jenkins,
	}, err
}

type Client struct {
	Address  string
	Username string
	Password string
	Jenkins  *gojenkins.Jenkins
}

func (c *Client) JobGetAll() ([]*gojenkins.Job, error) {
	return c.Jenkins.GetAllJobs(ctx)
}

func (c *Client) JobGet(jobName string) (*gojenkins.Job, error) {
	return c.Jenkins.GetJob(ctx, jobName)
}

func (c *Client) JobBuild(jobName string, params map[string]string) error {
	_, err := c.Jenkins.BuildJob(ctx, jobName, params)
	return err
}

func (c *Client) JobGetConfig(jobName string) (config string, err error) {
	job, err := c.Jenkins.GetJob(ctx, jobName)
	if err != nil {
		return "", err
	}
	return job.GetConfig(ctx)
}

func (c *Client) JobGetBuildAll(jobName string) ([]gojenkins.JobBuild, error) {
	return c.Jenkins.GetAllBuildIds(ctx, jobName)
}

func (c *Client) JobGetBuildLast(jobName string) (*gojenkins.Build, error) {
	job, err := c.JobGet(jobName)
	if err != nil {
		return nil, err
	}
	return job.GetLastBuild(ctx)
}
