// Code generated by protoc-gen-gohttp. DO NOT EDIT.
// source: model.proto

package __

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	io "io"
	ioutil "io/ioutil"
	mime "mime"
	http "net/http"
	strings "strings"
)

// RandomiseHTTPService is the server API for Randomise service.
type RandomiseHTTPService interface {
	Random(context.Context, *CommonRequest) (*CommonResponse, error)
}

// RandomiseHTTPConverter has a function to convert RandomiseHTTPService interface to http.HandlerFunc.
type RandomiseHTTPConverter struct {
	srv RandomiseHTTPService
}

// NewRandomiseHTTPConverter returns RandomiseHTTPConverter.
func NewRandomiseHTTPConverter(srv RandomiseHTTPService) *RandomiseHTTPConverter {
	return &RandomiseHTTPConverter{
		srv: srv,
	}
}

// Random returns RandomiseHTTPService interface's Random converted to http.HandlerFunc.
func (h *RandomiseHTTPConverter) Random(cb func(ctx context.Context, w http.ResponseWriter, r *http.Request, arg, ret proto.Message, err error), interceptors ...grpc.UnaryServerInterceptor) http.HandlerFunc {
	if cb == nil {
		cb = func(ctx context.Context, w http.ResponseWriter, r *http.Request, arg, ret proto.Message, err error) {
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				p := status.New(codes.Unknown, err.Error()).Proto()
				switch contentType, _, _ := mime.ParseMediaType(r.Header.Get("Content-Type")); contentType {
				case "application/protobuf", "application/x-protobuf":
					buf, err := proto.Marshal(p)
					if err != nil {
						return
					}
					if _, err := io.Copy(w, bytes.NewBuffer(buf)); err != nil {
						return
					}
				case "application/json":
					buf, err := protojson.Marshal(p)
					if err != nil {
						return
					}
					if _, err := io.Copy(w, bytes.NewBuffer(buf)); err != nil {
						return
					}
				default:
				}
			}
		}
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		contentType, _, _ := mime.ParseMediaType(r.Header.Get("Content-Type"))

		accepts := strings.Split(r.Header.Get("Accept"), ",")
		accept := accepts[0]
		if accept == "*/*" || accept == "" {
			if contentType != "" {
				accept = contentType
			} else {
				accept = "application/json"
			}
		}

		w.Header().Set("Content-Type", accept)

		arg := &CommonRequest{}
		if r.Method != http.MethodGet {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				cb(ctx, w, r, nil, nil, err)
				return
			}

			switch contentType {
			case "application/protobuf", "application/x-protobuf":
				if err := proto.Unmarshal(body, arg); err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
			case "application/json":
				if err := protojson.Unmarshal(body, arg); err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
			default:
				w.WriteHeader(http.StatusUnsupportedMediaType)
				_, err := fmt.Fprintf(w, "Unsupported Content-Type: %s", contentType)
				cb(ctx, w, r, nil, nil, err)
				return
			}
		}

		n := len(interceptors)
		chained := func(ctx context.Context, arg interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			chainer := func(currentInter grpc.UnaryServerInterceptor, currentHandler grpc.UnaryHandler) grpc.UnaryHandler {
				return func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
					return currentInter(currentCtx, currentReq, info, currentHandler)
				}
			}

			chainedHandler := handler
			for i := n - 1; i >= 0; i-- {
				chainedHandler = chainer(interceptors[i], chainedHandler)
			}
			return chainedHandler(ctx, arg)
		}

		info := &grpc.UnaryServerInfo{
			Server:     h.srv,
			FullMethod: "/entities.Randomise/Random",
		}

		handler := func(c context.Context, req interface{}) (interface{}, error) {
			return h.srv.Random(c, req.(*CommonRequest))
		}

		iret, err := chained(ctx, arg, info, handler)
		if err != nil {
			cb(ctx, w, r, arg, nil, err)
			return
		}

		ret, ok := iret.(*CommonResponse)
		if !ok {
			cb(ctx, w, r, arg, nil, fmt.Errorf("/entities.Randomise/Random: interceptors have not return CommonResponse"))
			return
		}

		switch accept {
		case "application/protobuf", "application/x-protobuf":
			buf, err := proto.Marshal(ret)
			if err != nil {
				cb(ctx, w, r, arg, ret, err)
				return
			}
			if _, err := io.Copy(w, bytes.NewBuffer(buf)); err != nil {
				cb(ctx, w, r, arg, ret, err)
				return
			}
		case "application/json":
			buf, err := protojson.Marshal(ret)
			if err != nil {
				cb(ctx, w, r, arg, ret, err)
				return
			}
			if _, err := io.Copy(w, bytes.NewBuffer(buf)); err != nil {
				cb(ctx, w, r, arg, ret, err)
				return
			}
		default:
			w.WriteHeader(http.StatusUnsupportedMediaType)
			_, err := fmt.Fprintf(w, "Unsupported Accept: %s", accept)
			cb(ctx, w, r, arg, ret, err)
			return
		}
		cb(ctx, w, r, arg, ret, nil)
	})
}

// RandomWithName returns Service name, Method name and RandomiseHTTPService interface's Random converted to http.HandlerFunc.
func (h *RandomiseHTTPConverter) RandomWithName(cb func(ctx context.Context, w http.ResponseWriter, r *http.Request, arg, ret proto.Message, err error), interceptors ...grpc.UnaryServerInterceptor) (string, string, http.HandlerFunc) {
	return "Randomise", "Random", h.Random(cb, interceptors...)
}
