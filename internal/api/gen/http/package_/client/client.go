// Code generated by goa v3.11.3, DO NOT EDIT.
//
// package client HTTP transport
//
// Command:
// $ goa gen github.com/artefactual-sdps/enduro/internal/api/design -o
// internal/api

package client

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the package service endpoint HTTP clients.
type Client struct {
	// MonitorRequest Doer is the HTTP client used to make requests to the
	// monitor_request endpoint.
	MonitorRequestDoer goahttp.Doer

	// Monitor Doer is the HTTP client used to make requests to the monitor
	// endpoint.
	MonitorDoer goahttp.Doer

	// List Doer is the HTTP client used to make requests to the list endpoint.
	ListDoer goahttp.Doer

	// Show Doer is the HTTP client used to make requests to the show endpoint.
	ShowDoer goahttp.Doer

	// PreservationActions Doer is the HTTP client used to make requests to the
	// preservation_actions endpoint.
	PreservationActionsDoer goahttp.Doer

	// Confirm Doer is the HTTP client used to make requests to the confirm
	// endpoint.
	ConfirmDoer goahttp.Doer

	// Reject Doer is the HTTP client used to make requests to the reject endpoint.
	RejectDoer goahttp.Doer

	// Move Doer is the HTTP client used to make requests to the move endpoint.
	MoveDoer goahttp.Doer

	// MoveStatus Doer is the HTTP client used to make requests to the move_status
	// endpoint.
	MoveStatusDoer goahttp.Doer

	// CORS Doer is the HTTP client used to make requests to the  endpoint.
	CORSDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme     string
	host       string
	encoder    func(*http.Request) goahttp.Encoder
	decoder    func(*http.Response) goahttp.Decoder
	dialer     goahttp.Dialer
	configurer *ConnConfigurer
}

// NewClient instantiates HTTP clients for all the package service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
	dialer goahttp.Dialer,
	cfn *ConnConfigurer,
) *Client {
	if cfn == nil {
		cfn = &ConnConfigurer{}
	}
	return &Client{
		MonitorRequestDoer:      doer,
		MonitorDoer:             doer,
		ListDoer:                doer,
		ShowDoer:                doer,
		PreservationActionsDoer: doer,
		ConfirmDoer:             doer,
		RejectDoer:              doer,
		MoveDoer:                doer,
		MoveStatusDoer:          doer,
		CORSDoer:                doer,
		RestoreResponseBody:     restoreBody,
		scheme:                  scheme,
		host:                    host,
		decoder:                 dec,
		encoder:                 enc,
		dialer:                  dialer,
		configurer:              cfn,
	}
}

// MonitorRequest returns an endpoint that makes HTTP requests to the package
// service monitor_request server.
func (c *Client) MonitorRequest() goa.Endpoint {
	var (
		encodeRequest  = EncodeMonitorRequestRequest(c.encoder)
		decodeResponse = DecodeMonitorRequestResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildMonitorRequestRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.MonitorRequestDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("package", "monitor_request", err)
		}
		return decodeResponse(resp)
	}
}

// Monitor returns an endpoint that makes HTTP requests to the package service
// monitor server.
func (c *Client) Monitor() goa.Endpoint {
	var (
		encodeRequest  = EncodeMonitorRequest(c.encoder)
		decodeResponse = DecodeMonitorResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildMonitorRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		conn, resp, err := c.dialer.DialContext(ctx, req.URL.String(), req.Header)
		if err != nil {
			if resp != nil {
				return decodeResponse(resp)
			}
			return nil, goahttp.ErrRequestError("package", "monitor", err)
		}
		if c.configurer.MonitorFn != nil {
			conn = c.configurer.MonitorFn(conn, cancel)
		}
		go func() {
			<-ctx.Done()
			conn.WriteControl(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, "client closing connection"),
				time.Now().Add(time.Second),
			)
			conn.Close()
		}()
		stream := &MonitorClientStream{conn: conn}
		return stream, nil
	}
}

// List returns an endpoint that makes HTTP requests to the package service
// list server.
func (c *Client) List() goa.Endpoint {
	var (
		encodeRequest  = EncodeListRequest(c.encoder)
		decodeResponse = DecodeListResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildListRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ListDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("package", "list", err)
		}
		return decodeResponse(resp)
	}
}

// Show returns an endpoint that makes HTTP requests to the package service
// show server.
func (c *Client) Show() goa.Endpoint {
	var (
		encodeRequest  = EncodeShowRequest(c.encoder)
		decodeResponse = DecodeShowResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildShowRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ShowDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("package", "show", err)
		}
		return decodeResponse(resp)
	}
}

// PreservationActions returns an endpoint that makes HTTP requests to the
// package service preservation_actions server.
func (c *Client) PreservationActions() goa.Endpoint {
	var (
		encodeRequest  = EncodePreservationActionsRequest(c.encoder)
		decodeResponse = DecodePreservationActionsResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildPreservationActionsRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.PreservationActionsDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("package", "preservation_actions", err)
		}
		return decodeResponse(resp)
	}
}

// Confirm returns an endpoint that makes HTTP requests to the package service
// confirm server.
func (c *Client) Confirm() goa.Endpoint {
	var (
		encodeRequest  = EncodeConfirmRequest(c.encoder)
		decodeResponse = DecodeConfirmResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildConfirmRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ConfirmDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("package", "confirm", err)
		}
		return decodeResponse(resp)
	}
}

// Reject returns an endpoint that makes HTTP requests to the package service
// reject server.
func (c *Client) Reject() goa.Endpoint {
	var (
		encodeRequest  = EncodeRejectRequest(c.encoder)
		decodeResponse = DecodeRejectResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildRejectRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RejectDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("package", "reject", err)
		}
		return decodeResponse(resp)
	}
}

// Move returns an endpoint that makes HTTP requests to the package service
// move server.
func (c *Client) Move() goa.Endpoint {
	var (
		encodeRequest  = EncodeMoveRequest(c.encoder)
		decodeResponse = DecodeMoveResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildMoveRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.MoveDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("package", "move", err)
		}
		return decodeResponse(resp)
	}
}

// MoveStatus returns an endpoint that makes HTTP requests to the package
// service move_status server.
func (c *Client) MoveStatus() goa.Endpoint {
	var (
		encodeRequest  = EncodeMoveStatusRequest(c.encoder)
		decodeResponse = DecodeMoveStatusResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildMoveStatusRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.MoveStatusDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("package", "move_status", err)
		}
		return decodeResponse(resp)
	}
}
