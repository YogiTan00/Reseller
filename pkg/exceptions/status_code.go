package exceptions

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func MapToGrpcStatusCode(err error) int {
	switch {
	case errors.Is(err, ErrInternalServer):
		return http.StatusInternalServerError
	case errors.Is(err, ErrUnauthorized):
		return http.StatusUnauthorized
	case errors.Is(err, ErrForbidden):
		return http.StatusForbidden
	default:
		return http.StatusBadRequest
	}
}

func MapToGrpcStatusError(err error) error {
	switch {
	case errors.Is(err, ErrInternalServer):
		return status.Error(codes.Internal, err.Error())
	case errors.Is(err, ErrUnauthorized):
		return status.Error(codes.Unauthenticated, err.Error())
	case errors.Is(err, ErrForbidden):
		return status.Error(codes.PermissionDenied, err.Error())
	default:
		return status.Error(codes.InvalidArgument, err.Error())
	}

}
