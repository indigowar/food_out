// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func encodeRetrieveDishByIDResponse(response RetrieveDishByIDRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *Dish:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *RetrieveDishByIDNotFound:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *RetrieveDishByIDInternalServerError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeRetrieveListOfRestaurantsResponse(response RetrieveListOfRestaurantsRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *RetrieveListOfRestaurantsOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *Error:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeRetrieveMenuByIDResponse(response RetrieveMenuByIDRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *Menu:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *RetrieveMenuByIDNotFound:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *RetrieveMenuByIDInternalServerError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeRetrieveMenusByRestaurantResponse(response RetrieveMenusByRestaurantRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *RetrieveMenusByRestaurantOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *RetrieveMenusByRestaurantNotFound:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *RetrieveMenusByRestaurantInternalServerError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeValidateRestaurantDishesResponse(response ValidateRestaurantDishesRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *ValidateRestaurantDishesOK:
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		return nil

	case *ValidateRestaurantDishesBadRequest:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *ValidateRestaurantDishesNotFound:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *ValidateRestaurantDishesInternalServerError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}
