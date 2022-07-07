package billing

import (
	"context"
	"fmt"

	billing "cloud.google.com/go/billing/apiv1"
	"google.golang.org/api/iterator"
	billingpb "google.golang.org/genproto/googleapis/cloud/billing/v1"
)

type GCP struct {
}

func (this *GCP) GetProjects() {
	err, billingaccountlist := GetAllBillingAccount()
	if err == nil {
		fmt.Println(err)
		return
	}
	for _, billingaccount := range billingaccountlist {
		err, projects := GetAllProject(billingaccount)
		if err == nil {
			fmt.Println(err)
			break
		}
		for _, project := range projects {
			fmt.Println(project)
		}
	}

}

func GetAllBillingAccount() (error, []string) {
	var result = make([]string, 0)
	ctx := context.Background()
	c, err := billing.NewCloudBillingClient(ctx)
	if err != nil {
		fmt.Println(err)
		return err, result
	}
	defer c.Close()
	req := &billingpb.ListBillingAccountsRequest{}
	it := c.ListBillingAccounts(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
			return err, result
		}
		result = append(result, resp.GetName())
	}
	return nil, result
}

func GetAllProject(billingaccount string) (error, []string) {
	var result = make([]string, 0)
	ctx := context.Background()
	c, err := billing.NewCloudBillingClient(ctx)
	if err != nil {
		fmt.Println(err)
		return err, result
	}
	defer c.Close()
	req := &billingpb.ListProjectBillingInfoRequest{
		Name: billingaccount,
	}
	it := c.ListProjectBillingInfo(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
			return err, result
		}
		result = append(result, resp.GetProjectId())
	}
	return nil, result
}
