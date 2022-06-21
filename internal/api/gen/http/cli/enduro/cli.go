// Code generated by goa v3.7.6, DO NOT EDIT.
//
// enduro HTTP client CLI support package
//
// Command:
// $ goa-v3.7.6 gen github.com/artefactual-labs/enduro/internal/api/design -o
// internal/api

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	batchc "github.com/artefactual-labs/enduro/internal/api/gen/http/batch/client"
	package_c "github.com/artefactual-labs/enduro/internal/api/gen/http/package_/client"
	storagec "github.com/artefactual-labs/enduro/internal/api/gen/http/storage/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `batch (submit|status|hints)
package (monitor|list|show|delete|cancel|retry|workflow|download|bulk|bulk-status|preservation-actions|accept|reject)
storage (submit|update|download)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` batch submit --body '{
      "completed_dir": "Quia commodi id.",
      "path": "Consequatur impedit vel exercitationem.",
      "retention_period": "Tempora eveniet eligendi."
   }'` + "\n" +
		os.Args[0] + ` package monitor` + "\n" +
		os.Args[0] + ` storage submit --body '{
      "name": "Omnis quisquam ad consequuntur."
   }' --aip-id "Nesciunt qui est provident fuga aut consequatur."` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
	dialer goahttp.Dialer,
	package_Configurer *package_c.ConnConfigurer,
) (goa.Endpoint, interface{}, error) {
	var (
		batchFlags = flag.NewFlagSet("batch", flag.ContinueOnError)

		batchSubmitFlags    = flag.NewFlagSet("submit", flag.ExitOnError)
		batchSubmitBodyFlag = batchSubmitFlags.String("body", "REQUIRED", "")

		batchStatusFlags = flag.NewFlagSet("status", flag.ExitOnError)

		batchHintsFlags = flag.NewFlagSet("hints", flag.ExitOnError)

		package_Flags = flag.NewFlagSet("package", flag.ContinueOnError)

		package_MonitorFlags = flag.NewFlagSet("monitor", flag.ExitOnError)

		package_ListFlags                   = flag.NewFlagSet("list", flag.ExitOnError)
		package_ListNameFlag                = package_ListFlags.String("name", "", "")
		package_ListAipIDFlag               = package_ListFlags.String("aip-id", "", "")
		package_ListEarliestCreatedTimeFlag = package_ListFlags.String("earliest-created-time", "", "")
		package_ListLatestCreatedTimeFlag   = package_ListFlags.String("latest-created-time", "", "")
		package_ListStatusFlag              = package_ListFlags.String("status", "", "")
		package_ListCursorFlag              = package_ListFlags.String("cursor", "", "")

		package_ShowFlags  = flag.NewFlagSet("show", flag.ExitOnError)
		package_ShowIDFlag = package_ShowFlags.String("id", "REQUIRED", "Identifier of package to show")

		package_DeleteFlags  = flag.NewFlagSet("delete", flag.ExitOnError)
		package_DeleteIDFlag = package_DeleteFlags.String("id", "REQUIRED", "Identifier of package to delete")

		package_CancelFlags  = flag.NewFlagSet("cancel", flag.ExitOnError)
		package_CancelIDFlag = package_CancelFlags.String("id", "REQUIRED", "Identifier of package to remove")

		package_RetryFlags  = flag.NewFlagSet("retry", flag.ExitOnError)
		package_RetryIDFlag = package_RetryFlags.String("id", "REQUIRED", "Identifier of package to retry")

		package_WorkflowFlags  = flag.NewFlagSet("workflow", flag.ExitOnError)
		package_WorkflowIDFlag = package_WorkflowFlags.String("id", "REQUIRED", "Identifier of package to look up")

		package_DownloadFlags  = flag.NewFlagSet("download", flag.ExitOnError)
		package_DownloadIDFlag = package_DownloadFlags.String("id", "REQUIRED", "Identifier of package to look up")

		package_BulkFlags    = flag.NewFlagSet("bulk", flag.ExitOnError)
		package_BulkBodyFlag = package_BulkFlags.String("body", "REQUIRED", "")

		package_BulkStatusFlags = flag.NewFlagSet("bulk-status", flag.ExitOnError)

		package_PreservationActionsFlags  = flag.NewFlagSet("preservation-actions", flag.ExitOnError)
		package_PreservationActionsIDFlag = package_PreservationActionsFlags.String("id", "REQUIRED", "Identifier of package to look up")

		package_AcceptFlags  = flag.NewFlagSet("accept", flag.ExitOnError)
		package_AcceptIDFlag = package_AcceptFlags.String("id", "REQUIRED", "Identifier of package to look up")

		package_RejectFlags  = flag.NewFlagSet("reject", flag.ExitOnError)
		package_RejectIDFlag = package_RejectFlags.String("id", "REQUIRED", "Identifier of package to look up")

		storageFlags = flag.NewFlagSet("storage", flag.ContinueOnError)

		storageSubmitFlags     = flag.NewFlagSet("submit", flag.ExitOnError)
		storageSubmitBodyFlag  = storageSubmitFlags.String("body", "REQUIRED", "")
		storageSubmitAipIDFlag = storageSubmitFlags.String("aip-id", "REQUIRED", "")

		storageUpdateFlags     = flag.NewFlagSet("update", flag.ExitOnError)
		storageUpdateAipIDFlag = storageUpdateFlags.String("aip-id", "REQUIRED", "")

		storageDownloadFlags     = flag.NewFlagSet("download", flag.ExitOnError)
		storageDownloadAipIDFlag = storageDownloadFlags.String("aip-id", "REQUIRED", "")
	)
	batchFlags.Usage = batchUsage
	batchSubmitFlags.Usage = batchSubmitUsage
	batchStatusFlags.Usage = batchStatusUsage
	batchHintsFlags.Usage = batchHintsUsage

	package_Flags.Usage = package_Usage
	package_MonitorFlags.Usage = package_MonitorUsage
	package_ListFlags.Usage = package_ListUsage
	package_ShowFlags.Usage = package_ShowUsage
	package_DeleteFlags.Usage = package_DeleteUsage
	package_CancelFlags.Usage = package_CancelUsage
	package_RetryFlags.Usage = package_RetryUsage
	package_WorkflowFlags.Usage = package_WorkflowUsage
	package_DownloadFlags.Usage = package_DownloadUsage
	package_BulkFlags.Usage = package_BulkUsage
	package_BulkStatusFlags.Usage = package_BulkStatusUsage
	package_PreservationActionsFlags.Usage = package_PreservationActionsUsage
	package_AcceptFlags.Usage = package_AcceptUsage
	package_RejectFlags.Usage = package_RejectUsage

	storageFlags.Usage = storageUsage
	storageSubmitFlags.Usage = storageSubmitUsage
	storageUpdateFlags.Usage = storageUpdateUsage
	storageDownloadFlags.Usage = storageDownloadUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "batch":
			svcf = batchFlags
		case "package":
			svcf = package_Flags
		case "storage":
			svcf = storageFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "batch":
			switch epn {
			case "submit":
				epf = batchSubmitFlags

			case "status":
				epf = batchStatusFlags

			case "hints":
				epf = batchHintsFlags

			}

		case "package":
			switch epn {
			case "monitor":
				epf = package_MonitorFlags

			case "list":
				epf = package_ListFlags

			case "show":
				epf = package_ShowFlags

			case "delete":
				epf = package_DeleteFlags

			case "cancel":
				epf = package_CancelFlags

			case "retry":
				epf = package_RetryFlags

			case "workflow":
				epf = package_WorkflowFlags

			case "download":
				epf = package_DownloadFlags

			case "bulk":
				epf = package_BulkFlags

			case "bulk-status":
				epf = package_BulkStatusFlags

			case "preservation-actions":
				epf = package_PreservationActionsFlags

			case "accept":
				epf = package_AcceptFlags

			case "reject":
				epf = package_RejectFlags

			}

		case "storage":
			switch epn {
			case "submit":
				epf = storageSubmitFlags

			case "update":
				epf = storageUpdateFlags

			case "download":
				epf = storageDownloadFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "batch":
			c := batchc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "submit":
				endpoint = c.Submit()
				data, err = batchc.BuildSubmitPayload(*batchSubmitBodyFlag)
			case "status":
				endpoint = c.Status()
				data = nil
			case "hints":
				endpoint = c.Hints()
				data = nil
			}
		case "package":
			c := package_c.NewClient(scheme, host, doer, enc, dec, restore, dialer, package_Configurer)
			switch epn {
			case "monitor":
				endpoint = c.Monitor()
				data = nil
			case "list":
				endpoint = c.List()
				data, err = package_c.BuildListPayload(*package_ListNameFlag, *package_ListAipIDFlag, *package_ListEarliestCreatedTimeFlag, *package_ListLatestCreatedTimeFlag, *package_ListStatusFlag, *package_ListCursorFlag)
			case "show":
				endpoint = c.Show()
				data, err = package_c.BuildShowPayload(*package_ShowIDFlag)
			case "delete":
				endpoint = c.Delete()
				data, err = package_c.BuildDeletePayload(*package_DeleteIDFlag)
			case "cancel":
				endpoint = c.Cancel()
				data, err = package_c.BuildCancelPayload(*package_CancelIDFlag)
			case "retry":
				endpoint = c.Retry()
				data, err = package_c.BuildRetryPayload(*package_RetryIDFlag)
			case "workflow":
				endpoint = c.Workflow()
				data, err = package_c.BuildWorkflowPayload(*package_WorkflowIDFlag)
			case "download":
				endpoint = c.Download()
				data, err = package_c.BuildDownloadPayload(*package_DownloadIDFlag)
			case "bulk":
				endpoint = c.Bulk()
				data, err = package_c.BuildBulkPayload(*package_BulkBodyFlag)
			case "bulk-status":
				endpoint = c.BulkStatus()
				data = nil
			case "preservation-actions":
				endpoint = c.PreservationActions()
				data, err = package_c.BuildPreservationActionsPayload(*package_PreservationActionsIDFlag)
			case "accept":
				endpoint = c.Accept()
				data, err = package_c.BuildAcceptPayload(*package_AcceptIDFlag)
			case "reject":
				endpoint = c.Reject()
				data, err = package_c.BuildRejectPayload(*package_RejectIDFlag)
			}
		case "storage":
			c := storagec.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "submit":
				endpoint = c.Submit()
				data, err = storagec.BuildSubmitPayload(*storageSubmitBodyFlag, *storageSubmitAipIDFlag)
			case "update":
				endpoint = c.Update()
				data, err = storagec.BuildUpdatePayload(*storageUpdateAipIDFlag)
			case "download":
				endpoint = c.Download()
				data, err = storagec.BuildDownloadPayload(*storageDownloadAipIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// batchUsage displays the usage of the batch command and its subcommands.
func batchUsage() {
	fmt.Fprintf(os.Stderr, `The batch service manages batches of packages.
Usage:
    %[1]s [globalflags] batch COMMAND [flags]

COMMAND:
    submit: Submit a new batch
    status: Retrieve status of current batch operation.
    hints: Retrieve form hints

Additional help:
    %[1]s batch COMMAND --help
`, os.Args[0])
}
func batchSubmitUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] batch submit -body JSON

Submit a new batch
    -body JSON: 

Example:
    %[1]s batch submit --body '{
      "completed_dir": "Quia commodi id.",
      "path": "Consequatur impedit vel exercitationem.",
      "retention_period": "Tempora eveniet eligendi."
   }'
`, os.Args[0])
}

func batchStatusUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] batch status

Retrieve status of current batch operation.

Example:
    %[1]s batch status
`, os.Args[0])
}

func batchHintsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] batch hints

Retrieve form hints

Example:
    %[1]s batch hints
`, os.Args[0])
}

// packageUsage displays the usage of the package command and its subcommands.
func package_Usage() {
	fmt.Fprintf(os.Stderr, `The package service manages packages being transferred to a3m.
Usage:
    %[1]s [globalflags] package COMMAND [flags]

COMMAND:
    monitor: Monitor implements monitor.
    list: List all stored packages
    show: Show package by ID
    delete: Delete package by ID
    cancel: Cancel package processing by ID
    retry: Retry package processing by ID
    workflow: Retrieve workflow status by ID
    download: Download package by ID
    bulk: Bulk operations (retry, cancel...).
    bulk-status: Retrieve status of current bulk operation.
    preservation-actions: List all preservation actions by ID
    accept: Signal the package has been reviewed and accepted
    reject: Signal the package has been reviewed and rejected

Additional help:
    %[1]s package COMMAND --help
`, os.Args[0])
}
func package_MonitorUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package monitor

Monitor implements monitor.

Example:
    %[1]s package monitor
`, os.Args[0])
}

func package_ListUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package list -name STRING -aip-id STRING -earliest-created-time STRING -latest-created-time STRING -status STRING -cursor STRING

List all stored packages
    -name STRING: 
    -aip-id STRING: 
    -earliest-created-time STRING: 
    -latest-created-time STRING: 
    -status STRING: 
    -cursor STRING: 

Example:
    %[1]s package list --name "Illo natus et itaque eaque." --aip-id "4CCDE767-7648-444F-D09F-4B4FFE4EB36B" --earliest-created-time "1993-03-03T14:33:12Z" --latest-created-time "2000-07-10T02:09:45Z" --status "error" --cursor "Dicta explicabo est quaerat incidunt."
`, os.Args[0])
}

func package_ShowUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package show -id UINT

Show package by ID
    -id UINT: Identifier of package to show

Example:
    %[1]s package show --id 3996361606384399983
`, os.Args[0])
}

func package_DeleteUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package delete -id UINT

Delete package by ID
    -id UINT: Identifier of package to delete

Example:
    %[1]s package delete --id 5829343242568083082
`, os.Args[0])
}

func package_CancelUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package cancel -id UINT

Cancel package processing by ID
    -id UINT: Identifier of package to remove

Example:
    %[1]s package cancel --id 11255156712036053872
`, os.Args[0])
}

func package_RetryUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package retry -id UINT

Retry package processing by ID
    -id UINT: Identifier of package to retry

Example:
    %[1]s package retry --id 4269239435361436165
`, os.Args[0])
}

func package_WorkflowUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package workflow -id UINT

Retrieve workflow status by ID
    -id UINT: Identifier of package to look up

Example:
    %[1]s package workflow --id 7057964250562282979
`, os.Args[0])
}

func package_DownloadUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package download -id UINT

Download package by ID
    -id UINT: Identifier of package to look up

Example:
    %[1]s package download --id 16208726300707896748
`, os.Args[0])
}

func package_BulkUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package bulk -body JSON

Bulk operations (retry, cancel...).
    -body JSON: 

Example:
    %[1]s package bulk --body '{
      "operation": "retry",
      "size": 13733910633092730650,
      "status": "done"
   }'
`, os.Args[0])
}

func package_BulkStatusUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package bulk-status

Retrieve status of current bulk operation.

Example:
    %[1]s package bulk-status
`, os.Args[0])
}

func package_PreservationActionsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package preservation-actions -id UINT

List all preservation actions by ID
    -id UINT: Identifier of package to look up

Example:
    %[1]s package preservation-actions --id 11970900180965159798
`, os.Args[0])
}

func package_AcceptUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package accept -id UINT

Signal the package has been reviewed and accepted
    -id UINT: Identifier of package to look up

Example:
    %[1]s package accept --id 15253041435799747880
`, os.Args[0])
}

func package_RejectUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package reject -id UINT

Signal the package has been reviewed and rejected
    -id UINT: Identifier of package to look up

Example:
    %[1]s package reject --id 13110967601727961339
`, os.Args[0])
}

// storageUsage displays the usage of the storage command and its subcommands.
func storageUsage() {
	fmt.Fprintf(os.Stderr, `The storage service manages the storage of packages.
Usage:
    %[1]s [globalflags] storage COMMAND [flags]

COMMAND:
    submit: Start the submission of a package
    update: Signal the storage service that an upload is complete
    download: Download package by AIPID

Additional help:
    %[1]s storage COMMAND --help
`, os.Args[0])
}
func storageSubmitUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage submit -body JSON -aip-id STRING

Start the submission of a package
    -body JSON: 
    -aip-id STRING: 

Example:
    %[1]s storage submit --body '{
      "name": "Omnis quisquam ad consequuntur."
   }' --aip-id "Nesciunt qui est provident fuga aut consequatur."
`, os.Args[0])
}

func storageUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage update -aip-id STRING

Signal the storage service that an upload is complete
    -aip-id STRING: 

Example:
    %[1]s storage update --aip-id "Ducimus totam atque pariatur."
`, os.Args[0])
}

func storageDownloadUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage download -aip-id STRING

Download package by AIPID
    -aip-id STRING: 

Example:
    %[1]s storage download --aip-id "Modi maiores sit veniam."
`, os.Args[0])
}
