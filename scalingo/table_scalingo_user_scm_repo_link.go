package scalingo

import (
	"context"

	"github.com/Scalingo/go-scalingo/v9"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableScalingoUserScmRepoLink() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_user_scm_repo_link",
		Description: "List all the SCM links associated to your account.",
		List: &plugin.ListConfig{
			Hydrate: listUserScmRepoLink,
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns:           scalingoSCMRepoLinkColumns(),
	}
}

func listUserScmRepoLink(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_user_scm_repo_link.listUserScmRepoLink", "connection_error", err)
		return nil, err
	}
	opts := scalingo.PaginationOpts{Page: 1, PerPage: 50}
	for {
		scmRepoLinks, pagination, err := client.SCMRepoLinkList(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("scalingo_user_scm_repo_link.listUserScmRepoLink", err)
			return nil, err
		}
		for _, scmRepoLink := range scmRepoLinks {
			d.StreamListItem(ctx, scmRepoLink)
		}
		if pagination.NextPage == 0 {
			break
		}
		opts.Page = pagination.NextPage
	}
	return nil, nil
}
