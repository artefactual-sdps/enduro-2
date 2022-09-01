// Code generated by goa v3.8.4, DO NOT EDIT.
//
// enduro HTTP client CLI support package
//
// Command:
// $ goa-v3.8.4 gen github.com/artefactual-sdps/enduro/internal/api/design -o
// internal/api

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	package_c "github.com/artefactual-sdps/enduro/internal/api/gen/http/package_/client"
	storagec "github.com/artefactual-sdps/enduro/internal/api/gen/http/storage/client"
	uploadc "github.com/artefactual-sdps/enduro/internal/api/gen/http/upload/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `package (monitor|list|show|preservation-actions|confirm|reject|move|move-status)
storage (submit|update|download|locations|add-location|move|move-status|reject|show|show-location|location-packages)
upload upload
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` package monitor` + "\n" +
		os.Args[0] + ` storage submit --body '{
      "name": "Est quia et praesentium a autem."
   }' --aip-id "0191B0B6-8FCE-0417-5FE7-F3E5947851AE"` + "\n" +
		os.Args[0] + ` upload upload --content-type "multipart/󋸺񫇖󬈖𧯒󥡮; boundary=���" --stream "goa.png"` + "\n" +
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
		package_Flags = flag.NewFlagSet("package", flag.ContinueOnError)

		package_MonitorFlags = flag.NewFlagSet("monitor", flag.ExitOnError)

		package_ListFlags                   = flag.NewFlagSet("list", flag.ExitOnError)
		package_ListNameFlag                = package_ListFlags.String("name", "", "")
		package_ListAipIDFlag               = package_ListFlags.String("aip-id", "", "")
		package_ListEarliestCreatedTimeFlag = package_ListFlags.String("earliest-created-time", "", "")
		package_ListLatestCreatedTimeFlag   = package_ListFlags.String("latest-created-time", "", "")
		package_ListLocationIDFlag          = package_ListFlags.String("location-id", "", "")
		package_ListStatusFlag              = package_ListFlags.String("status", "", "")
		package_ListCursorFlag              = package_ListFlags.String("cursor", "", "")

		package_ShowFlags  = flag.NewFlagSet("show", flag.ExitOnError)
		package_ShowIDFlag = package_ShowFlags.String("id", "REQUIRED", "Identifier of package to show")

		package_PreservationActionsFlags  = flag.NewFlagSet("preservation-actions", flag.ExitOnError)
		package_PreservationActionsIDFlag = package_PreservationActionsFlags.String("id", "REQUIRED", "Identifier of package to look up")

		package_ConfirmFlags    = flag.NewFlagSet("confirm", flag.ExitOnError)
		package_ConfirmBodyFlag = package_ConfirmFlags.String("body", "REQUIRED", "")
		package_ConfirmIDFlag   = package_ConfirmFlags.String("id", "REQUIRED", "Identifier of package to look up")

		package_RejectFlags  = flag.NewFlagSet("reject", flag.ExitOnError)
		package_RejectIDFlag = package_RejectFlags.String("id", "REQUIRED", "Identifier of package to look up")

		package_MoveFlags    = flag.NewFlagSet("move", flag.ExitOnError)
		package_MoveBodyFlag = package_MoveFlags.String("body", "REQUIRED", "")
		package_MoveIDFlag   = package_MoveFlags.String("id", "REQUIRED", "Identifier of package to move")

		package_MoveStatusFlags  = flag.NewFlagSet("move-status", flag.ExitOnError)
		package_MoveStatusIDFlag = package_MoveStatusFlags.String("id", "REQUIRED", "Identifier of package to move")

		storageFlags = flag.NewFlagSet("storage", flag.ContinueOnError)

		storageSubmitFlags     = flag.NewFlagSet("submit", flag.ExitOnError)
		storageSubmitBodyFlag  = storageSubmitFlags.String("body", "REQUIRED", "")
		storageSubmitAipIDFlag = storageSubmitFlags.String("aip-id", "REQUIRED", "")

		storageUpdateFlags     = flag.NewFlagSet("update", flag.ExitOnError)
		storageUpdateAipIDFlag = storageUpdateFlags.String("aip-id", "REQUIRED", "")

		storageDownloadFlags     = flag.NewFlagSet("download", flag.ExitOnError)
		storageDownloadAipIDFlag = storageDownloadFlags.String("aip-id", "REQUIRED", "")

		storageLocationsFlags = flag.NewFlagSet("locations", flag.ExitOnError)

		storageAddLocationFlags    = flag.NewFlagSet("add-location", flag.ExitOnError)
		storageAddLocationBodyFlag = storageAddLocationFlags.String("body", "REQUIRED", "")

		storageMoveFlags     = flag.NewFlagSet("move", flag.ExitOnError)
		storageMoveBodyFlag  = storageMoveFlags.String("body", "REQUIRED", "")
		storageMoveAipIDFlag = storageMoveFlags.String("aip-id", "REQUIRED", "")

		storageMoveStatusFlags     = flag.NewFlagSet("move-status", flag.ExitOnError)
		storageMoveStatusAipIDFlag = storageMoveStatusFlags.String("aip-id", "REQUIRED", "")

		storageRejectFlags     = flag.NewFlagSet("reject", flag.ExitOnError)
		storageRejectAipIDFlag = storageRejectFlags.String("aip-id", "REQUIRED", "")

		storageShowFlags     = flag.NewFlagSet("show", flag.ExitOnError)
		storageShowAipIDFlag = storageShowFlags.String("aip-id", "REQUIRED", "")

		storageShowLocationFlags    = flag.NewFlagSet("show-location", flag.ExitOnError)
		storageShowLocationUUIDFlag = storageShowLocationFlags.String("uuid", "REQUIRED", "")

		storageLocationPackagesFlags    = flag.NewFlagSet("location-packages", flag.ExitOnError)
		storageLocationPackagesUUIDFlag = storageLocationPackagesFlags.String("uuid", "REQUIRED", "")

		uploadFlags = flag.NewFlagSet("upload", flag.ContinueOnError)

		uploadUploadFlags           = flag.NewFlagSet("upload", flag.ExitOnError)
		uploadUploadContentTypeFlag = uploadUploadFlags.String("content-type", "multipart/form-data; boundary=goa", "")
		uploadUploadStreamFlag      = uploadUploadFlags.String("stream", "REQUIRED", "path to file containing the streamed request body")
	)
	package_Flags.Usage = package_Usage
	package_MonitorFlags.Usage = package_MonitorUsage
	package_ListFlags.Usage = package_ListUsage
	package_ShowFlags.Usage = package_ShowUsage
	package_PreservationActionsFlags.Usage = package_PreservationActionsUsage
	package_ConfirmFlags.Usage = package_ConfirmUsage
	package_RejectFlags.Usage = package_RejectUsage
	package_MoveFlags.Usage = package_MoveUsage
	package_MoveStatusFlags.Usage = package_MoveStatusUsage

	storageFlags.Usage = storageUsage
	storageSubmitFlags.Usage = storageSubmitUsage
	storageUpdateFlags.Usage = storageUpdateUsage
	storageDownloadFlags.Usage = storageDownloadUsage
	storageLocationsFlags.Usage = storageLocationsUsage
	storageAddLocationFlags.Usage = storageAddLocationUsage
	storageMoveFlags.Usage = storageMoveUsage
	storageMoveStatusFlags.Usage = storageMoveStatusUsage
	storageRejectFlags.Usage = storageRejectUsage
	storageShowFlags.Usage = storageShowUsage
	storageShowLocationFlags.Usage = storageShowLocationUsage
	storageLocationPackagesFlags.Usage = storageLocationPackagesUsage

	uploadFlags.Usage = uploadUsage
	uploadUploadFlags.Usage = uploadUploadUsage

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
		case "package":
			svcf = package_Flags
		case "storage":
			svcf = storageFlags
		case "upload":
			svcf = uploadFlags
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
		case "package":
			switch epn {
			case "monitor":
				epf = package_MonitorFlags

			case "list":
				epf = package_ListFlags

			case "show":
				epf = package_ShowFlags

			case "preservation-actions":
				epf = package_PreservationActionsFlags

			case "confirm":
				epf = package_ConfirmFlags

			case "reject":
				epf = package_RejectFlags

			case "move":
				epf = package_MoveFlags

			case "move-status":
				epf = package_MoveStatusFlags

			}

		case "storage":
			switch epn {
			case "submit":
				epf = storageSubmitFlags

			case "update":
				epf = storageUpdateFlags

			case "download":
				epf = storageDownloadFlags

			case "locations":
				epf = storageLocationsFlags

			case "add-location":
				epf = storageAddLocationFlags

			case "move":
				epf = storageMoveFlags

			case "move-status":
				epf = storageMoveStatusFlags

			case "reject":
				epf = storageRejectFlags

			case "show":
				epf = storageShowFlags

			case "show-location":
				epf = storageShowLocationFlags

			case "location-packages":
				epf = storageLocationPackagesFlags

			}

		case "upload":
			switch epn {
			case "upload":
				epf = uploadUploadFlags

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
		case "package":
			c := package_c.NewClient(scheme, host, doer, enc, dec, restore, dialer, package_Configurer)
			switch epn {
			case "monitor":
				endpoint = c.Monitor()
				data = nil
			case "list":
				endpoint = c.List()
				data, err = package_c.BuildListPayload(*package_ListNameFlag, *package_ListAipIDFlag, *package_ListEarliestCreatedTimeFlag, *package_ListLatestCreatedTimeFlag, *package_ListLocationIDFlag, *package_ListStatusFlag, *package_ListCursorFlag)
			case "show":
				endpoint = c.Show()
				data, err = package_c.BuildShowPayload(*package_ShowIDFlag)
			case "preservation-actions":
				endpoint = c.PreservationActions()
				data, err = package_c.BuildPreservationActionsPayload(*package_PreservationActionsIDFlag)
			case "confirm":
				endpoint = c.Confirm()
				data, err = package_c.BuildConfirmPayload(*package_ConfirmBodyFlag, *package_ConfirmIDFlag)
			case "reject":
				endpoint = c.Reject()
				data, err = package_c.BuildRejectPayload(*package_RejectIDFlag)
			case "move":
				endpoint = c.Move()
				data, err = package_c.BuildMovePayload(*package_MoveBodyFlag, *package_MoveIDFlag)
			case "move-status":
				endpoint = c.MoveStatus()
				data, err = package_c.BuildMoveStatusPayload(*package_MoveStatusIDFlag)
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
			case "locations":
				endpoint = c.Locations()
				data = nil
			case "add-location":
				endpoint = c.AddLocation()
				data, err = storagec.BuildAddLocationPayload(*storageAddLocationBodyFlag)
			case "move":
				endpoint = c.Move()
				data, err = storagec.BuildMovePayload(*storageMoveBodyFlag, *storageMoveAipIDFlag)
			case "move-status":
				endpoint = c.MoveStatus()
				data, err = storagec.BuildMoveStatusPayload(*storageMoveStatusAipIDFlag)
			case "reject":
				endpoint = c.Reject()
				data, err = storagec.BuildRejectPayload(*storageRejectAipIDFlag)
			case "show":
				endpoint = c.Show()
				data, err = storagec.BuildShowPayload(*storageShowAipIDFlag)
			case "show-location":
				endpoint = c.ShowLocation()
				data, err = storagec.BuildShowLocationPayload(*storageShowLocationUUIDFlag)
			case "location-packages":
				endpoint = c.LocationPackages()
				data, err = storagec.BuildLocationPackagesPayload(*storageLocationPackagesUUIDFlag)
			}
		case "upload":
			c := uploadc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "upload":
				endpoint = c.Upload()
				data, err = uploadc.BuildUploadPayload(*uploadUploadContentTypeFlag)
				if err == nil {
					data, err = uploadc.BuildUploadStreamPayload(data, *uploadUploadStreamFlag)
				}
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
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
    preservation-actions: List all preservation actions by ID
    confirm: Signal the package has been reviewed and accepted
    reject: Signal the package has been reviewed and rejected
    move: Move a package to a permanent storage location
    move-status: Retrieve the status of a permanent storage location move of the package

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
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package list -name STRING -aip-id STRING -earliest-created-time STRING -latest-created-time STRING -location-id STRING -status STRING -cursor STRING

List all stored packages
    -name STRING: 
    -aip-id STRING: 
    -earliest-created-time STRING: 
    -latest-created-time STRING: 
    -location-id STRING: 
    -status STRING: 
    -cursor STRING: 

Example:
    %[1]s package list --name "Voluptatem iusto consequatur qui aut ducimus totam." --aip-id "AB360986-C49D-9ED1-C841-8564F85C15A9" --earliest-created-time "1971-09-30T17:15:51Z" --latest-created-time "1988-06-14T22:34:25Z" --location-id "9609F161-371C-A63E-3657-D3B6F178B07F" --status "new" --cursor "Odit omnis itaque soluta sed et modi."
`, os.Args[0])
}

func package_ShowUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package show -id UINT

Show package by ID
    -id UINT: Identifier of package to show

Example:
    %[1]s package show --id 18178393672110633282
`, os.Args[0])
}

func package_PreservationActionsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package preservation-actions -id UINT

List all preservation actions by ID
    -id UINT: Identifier of package to look up

Example:
    %[1]s package preservation-actions --id 5007678243935735816
`, os.Args[0])
}

func package_ConfirmUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package confirm -body JSON -id UINT

Signal the package has been reviewed and accepted
    -body JSON: 
    -id UINT: Identifier of package to look up

Example:
    %[1]s package confirm --body '{
      "location_id": "Ut quas."
   }' --id 6413953064518636787
`, os.Args[0])
}

func package_RejectUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package reject -id UINT

Signal the package has been reviewed and rejected
    -id UINT: Identifier of package to look up

Example:
    %[1]s package reject --id 14757473849653299790
`, os.Args[0])
}

func package_MoveUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package move -body JSON -id UINT

Move a package to a permanent storage location
    -body JSON: 
    -id UINT: Identifier of package to move

Example:
    %[1]s package move --body '{
      "location_id": "Temporibus iusto et."
   }' --id 5413568045811429472
`, os.Args[0])
}

func package_MoveStatusUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package move-status -id UINT

Retrieve the status of a permanent storage location move of the package
    -id UINT: Identifier of package to move

Example:
    %[1]s package move-status --id 14030264248050019902
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
    locations: List locations
    add-location: Add a storage location
    move: Move a package to a permanent storage location
    move-status: Retrieve the status of a permanent storage location move of the package
    reject: Reject a package
    show: Show package by AIPID
    show-location: Show location by UUID
    location-packages: List all the packages stored in the location with UUID

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
      "name": "Est quia et praesentium a autem."
   }' --aip-id "0191B0B6-8FCE-0417-5FE7-F3E5947851AE"
`, os.Args[0])
}

func storageUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage update -aip-id STRING

Signal the storage service that an upload is complete
    -aip-id STRING: 

Example:
    %[1]s storage update --aip-id "C353393C-52D3-E5B1-A8C2-BA49F9BA7DA5"
`, os.Args[0])
}

func storageDownloadUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage download -aip-id STRING

Download package by AIPID
    -aip-id STRING: 

Example:
    %[1]s storage download --aip-id "F55E953C-3F94-1E57-76B3-19E312C289C8"
`, os.Args[0])
}

func storageLocationsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage locations

List locations

Example:
    %[1]s storage locations
`, os.Args[0])
}

func storageAddLocationUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage add-location -body JSON

Add a storage location
    -body JSON: 

Example:
    %[1]s storage add-location --body '{
      "config": {
         "Type": "s3",
         "Value": "\"JSON\""
      },
      "description": "Est sed.",
      "name": "Quam aut sit quo.",
      "purpose": "aip_store",
      "source": "minio"
   }'
`, os.Args[0])
}

func storageMoveUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage move -body JSON -aip-id STRING

Move a package to a permanent storage location
    -body JSON: 
    -aip-id STRING: 

Example:
    %[1]s storage move --body '{
      "location_id": "Dignissimos deserunt autem."
   }' --aip-id "05AF9F85-3F51-2CC5-AA06-E15963A2BFCC"
`, os.Args[0])
}

func storageMoveStatusUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage move-status -aip-id STRING

Retrieve the status of a permanent storage location move of the package
    -aip-id STRING: 

Example:
    %[1]s storage move-status --aip-id "79D69965-CB4F-B5AD-C251-2DDAE0E285F8"
`, os.Args[0])
}

func storageRejectUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage reject -aip-id STRING

Reject a package
    -aip-id STRING: 

Example:
    %[1]s storage reject --aip-id "F2B3804F-8F25-5448-962C-F0EBE08C813C"
`, os.Args[0])
}

func storageShowUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage show -aip-id STRING

Show package by AIPID
    -aip-id STRING: 

Example:
    %[1]s storage show --aip-id "8592D8CE-5388-87CA-26CB-3AD1899F8CA9"
`, os.Args[0])
}

func storageShowLocationUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage show-location -uuid STRING

Show location by UUID
    -uuid STRING: 

Example:
    %[1]s storage show-location --uuid "17DDC9D3-F205-F8F5-747D-4A2F0DDDD04A"
`, os.Args[0])
}

func storageLocationPackagesUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage location-packages -uuid STRING

List all the packages stored in the location with UUID
    -uuid STRING: 

Example:
    %[1]s storage location-packages --uuid "F2585014-8A98-729F-BBAB-7664BE8C6476"
`, os.Args[0])
}

// uploadUsage displays the usage of the upload command and its subcommands.
func uploadUsage() {
	fmt.Fprintf(os.Stderr, `The upload service handles file submissions to the SIPs bucket.
Usage:
    %[1]s [globalflags] upload COMMAND [flags]

COMMAND:
    upload: Upload implements upload.

Additional help:
    %[1]s upload COMMAND --help
`, os.Args[0])
}
func uploadUploadUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] upload upload -content-type STRING -stream STRING

Upload implements upload.
    -content-type STRING: 
    -stream STRING: path to file containing the streamed request body

Example:
    %[1]s upload upload --content-type "multipart/󋸺񫇖󬈖𧯒󥡮; boundary=���" --stream "goa.png"
`, os.Args[0])
}
