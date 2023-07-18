package project

import (
	"fmt"

	projectapi "github.com/CircleCI-Public/circleci-cli/api/project"
	"github.com/CircleCI-Public/circleci-cli/cmd/validator"
	"github.com/spf13/cobra"
)

func getWebhook() error {
	return nil
}

func listWebhook(cmd *cobra.Command, client projectapi.ProjectClient, vcsType, orgName, projName string) error {
	
	projectInfo, err := client.ProjectInfo(vcsType, orgName, projName)
	if err != nil {
		return err
	}
	projectId := projectInfo.Id

	fmt.Print(projectId)
	
	client.ListProjectWebhook(projectId)

	return nil
}

func addWebhook() error {
	return nil
}

func updateWebhook() error {
	return nil
}

func deleteWebhook() error {
	return nil
}

func newWebhookCommand(ops *projectOpts, preRunE validator.Validator) *cobra.Command {
	cmd := &cobra.Command{
		Use: "webhook",
		Short: "Operate on webhook of project",
	}

	listWebhook := &cobra.Command{
		Short: "List project webhooks",
		Use: "list <vcs-type> <org-name> <project-name>",
		PreRunE: preRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listWebhook(cmd, ops.projectClient, args[0], args[1], args[2])
		},
		Args: cobra.ExactArgs(3),
	}

	getWebhook := &cobra.Command{
		Short: "Get project webhook",
		Use: "get",
		PreRunE: preRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			return getWebhook()
		},
	}

	cmd.AddCommand(listWebhook)
	cmd.AddCommand(getWebhook)
	return cmd
}
