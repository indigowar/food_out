// Code generated by ogen, DO NOT EDIT.

package api

import (
	"math/bits"
	"strconv"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"

	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/validate"
)

// Encode implements json.Marshaler.
func (s *Dish) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Dish) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("id")
		json.EncodeUUID(e, s.ID)
	}
	{
		e.FieldStart("name")
		e.Str(s.Name)
	}
	{
		e.FieldStart("image")
		json.EncodeURI(e, s.Image)
	}
	{
		e.FieldStart("price")
		e.Float64(s.Price)
	}
}

var jsonFieldsNameOfDish = [4]string{
	0: "id",
	1: "name",
	2: "image",
	3: "price",
}

// Decode decodes Dish from json.
func (s *Dish) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Dish to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := json.DecodeUUID(d)
				s.ID = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "name":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Name = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		case "image":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := json.DecodeURI(d)
				s.Image = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"image\"")
			}
		case "price":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := d.Float64()
				s.Price = float64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"price\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Dish")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00001111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfDish) {
					name = jsonFieldsNameOfDish[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Dish) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Dish) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Error) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Error) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("code")
		e.Float64(s.Code)
	}
	{
		e.FieldStart("message")
		e.Str(s.Message)
	}
}

var jsonFieldsNameOfError = [2]string{
	0: "code",
	1: "message",
}

// Decode decodes Error from json.
func (s *Error) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Error to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "code":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Float64()
				s.Code = float64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"code\"")
			}
		case "message":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Message = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Error")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfError) {
					name = jsonFieldsNameOfError[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Error) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Error) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Menu) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Menu) encodeFields(e *jx.Encoder) {
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.Restaurant.Set {
			e.FieldStart("restaurant")
			s.Restaurant.Encode(e)
		}
	}
	{
		if s.Image.Set {
			e.FieldStart("image")
			s.Image.Encode(e)
		}
	}
	{
		if s.Dishes != nil {
			e.FieldStart("dishes")
			e.ArrStart()
			for _, elem := range s.Dishes {
				elem.Encode(e)
			}
			e.ArrEnd()
		}
	}
}

var jsonFieldsNameOfMenu = [4]string{
	0: "id",
	1: "restaurant",
	2: "image",
	3: "dishes",
}

// Decode decodes Menu from json.
func (s *Menu) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Menu to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "restaurant":
			if err := func() error {
				s.Restaurant.Reset()
				if err := s.Restaurant.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"restaurant\"")
			}
		case "image":
			if err := func() error {
				s.Image.Reset()
				if err := s.Image.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"image\"")
			}
		case "dishes":
			if err := func() error {
				s.Dishes = make([]Dish, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem Dish
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Dishes = append(s.Dishes, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"dishes\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Menu")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Menu) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Menu) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes url.URL as json.
func (o OptURI) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	json.EncodeURI(e, o.Value)
}

// Decode decodes url.URL from json.
func (o *OptURI) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptURI to nil")
	}
	o.Set = true
	v, err := json.DecodeURI(d)
	if err != nil {
		return err
	}
	o.Value = v
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptURI) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptURI) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes uuid.UUID as json.
func (o OptUUID) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	json.EncodeUUID(e, o.Value)
}

// Decode decodes uuid.UUID from json.
func (o *OptUUID) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptUUID to nil")
	}
	o.Set = true
	v, err := json.DecodeUUID(d)
	if err != nil {
		return err
	}
	o.Value = v
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptUUID) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptUUID) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes RetrieveDishByIDInternalServerError as json.
func (s *RetrieveDishByIDInternalServerError) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes RetrieveDishByIDInternalServerError from json.
func (s *RetrieveDishByIDInternalServerError) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RetrieveDishByIDInternalServerError to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = RetrieveDishByIDInternalServerError(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *RetrieveDishByIDInternalServerError) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *RetrieveDishByIDInternalServerError) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes RetrieveDishByIDNotFound as json.
func (s *RetrieveDishByIDNotFound) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes RetrieveDishByIDNotFound from json.
func (s *RetrieveDishByIDNotFound) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RetrieveDishByIDNotFound to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = RetrieveDishByIDNotFound(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *RetrieveDishByIDNotFound) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *RetrieveDishByIDNotFound) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes RetrieveDishesByMenuIdInternalServerError as json.
func (s *RetrieveDishesByMenuIdInternalServerError) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes RetrieveDishesByMenuIdInternalServerError from json.
func (s *RetrieveDishesByMenuIdInternalServerError) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RetrieveDishesByMenuIdInternalServerError to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = RetrieveDishesByMenuIdInternalServerError(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *RetrieveDishesByMenuIdInternalServerError) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *RetrieveDishesByMenuIdInternalServerError) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes RetrieveDishesByMenuIdNotFound as json.
func (s *RetrieveDishesByMenuIdNotFound) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes RetrieveDishesByMenuIdNotFound from json.
func (s *RetrieveDishesByMenuIdNotFound) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RetrieveDishesByMenuIdNotFound to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = RetrieveDishesByMenuIdNotFound(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *RetrieveDishesByMenuIdNotFound) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *RetrieveDishesByMenuIdNotFound) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes RetrieveDishesByMenuIdOKApplicationJSON as json.
func (s RetrieveDishesByMenuIdOKApplicationJSON) Encode(e *jx.Encoder) {
	unwrapped := []Dish(s)

	e.ArrStart()
	for _, elem := range unwrapped {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes RetrieveDishesByMenuIdOKApplicationJSON from json.
func (s *RetrieveDishesByMenuIdOKApplicationJSON) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RetrieveDishesByMenuIdOKApplicationJSON to nil")
	}
	var unwrapped []Dish
	if err := func() error {
		unwrapped = make([]Dish, 0)
		if err := d.Arr(func(d *jx.Decoder) error {
			var elem Dish
			if err := elem.Decode(d); err != nil {
				return err
			}
			unwrapped = append(unwrapped, elem)
			return nil
		}); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = RetrieveDishesByMenuIdOKApplicationJSON(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s RetrieveDishesByMenuIdOKApplicationJSON) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *RetrieveDishesByMenuIdOKApplicationJSON) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes RetrieveListOfRestaurantsOKApplicationJSON as json.
func (s RetrieveListOfRestaurantsOKApplicationJSON) Encode(e *jx.Encoder) {
	unwrapped := []uuid.UUID(s)

	e.ArrStart()
	for _, elem := range unwrapped {
		json.EncodeUUID(e, elem)
	}
	e.ArrEnd()
}

// Decode decodes RetrieveListOfRestaurantsOKApplicationJSON from json.
func (s *RetrieveListOfRestaurantsOKApplicationJSON) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RetrieveListOfRestaurantsOKApplicationJSON to nil")
	}
	var unwrapped []uuid.UUID
	if err := func() error {
		unwrapped = make([]uuid.UUID, 0)
		if err := d.Arr(func(d *jx.Decoder) error {
			var elem uuid.UUID
			v, err := json.DecodeUUID(d)
			elem = v
			if err != nil {
				return err
			}
			unwrapped = append(unwrapped, elem)
			return nil
		}); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = RetrieveListOfRestaurantsOKApplicationJSON(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s RetrieveListOfRestaurantsOKApplicationJSON) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *RetrieveListOfRestaurantsOKApplicationJSON) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes RetrieveMenuByIDInternalServerError as json.
func (s *RetrieveMenuByIDInternalServerError) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes RetrieveMenuByIDInternalServerError from json.
func (s *RetrieveMenuByIDInternalServerError) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RetrieveMenuByIDInternalServerError to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = RetrieveMenuByIDInternalServerError(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *RetrieveMenuByIDInternalServerError) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *RetrieveMenuByIDInternalServerError) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes RetrieveMenuByIDNotFound as json.
func (s *RetrieveMenuByIDNotFound) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes RetrieveMenuByIDNotFound from json.
func (s *RetrieveMenuByIDNotFound) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RetrieveMenuByIDNotFound to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = RetrieveMenuByIDNotFound(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *RetrieveMenuByIDNotFound) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *RetrieveMenuByIDNotFound) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes RetrieveMenusByRestaurantInternalServerError as json.
func (s *RetrieveMenusByRestaurantInternalServerError) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes RetrieveMenusByRestaurantInternalServerError from json.
func (s *RetrieveMenusByRestaurantInternalServerError) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RetrieveMenusByRestaurantInternalServerError to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = RetrieveMenusByRestaurantInternalServerError(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *RetrieveMenusByRestaurantInternalServerError) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *RetrieveMenusByRestaurantInternalServerError) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes RetrieveMenusByRestaurantNotFound as json.
func (s *RetrieveMenusByRestaurantNotFound) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes RetrieveMenusByRestaurantNotFound from json.
func (s *RetrieveMenusByRestaurantNotFound) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RetrieveMenusByRestaurantNotFound to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = RetrieveMenusByRestaurantNotFound(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *RetrieveMenusByRestaurantNotFound) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *RetrieveMenusByRestaurantNotFound) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes RetrieveMenusByRestaurantOKApplicationJSON as json.
func (s RetrieveMenusByRestaurantOKApplicationJSON) Encode(e *jx.Encoder) {
	unwrapped := []Menu(s)

	e.ArrStart()
	for _, elem := range unwrapped {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes RetrieveMenusByRestaurantOKApplicationJSON from json.
func (s *RetrieveMenusByRestaurantOKApplicationJSON) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RetrieveMenusByRestaurantOKApplicationJSON to nil")
	}
	var unwrapped []Menu
	if err := func() error {
		unwrapped = make([]Menu, 0)
		if err := d.Arr(func(d *jx.Decoder) error {
			var elem Menu
			if err := elem.Decode(d); err != nil {
				return err
			}
			unwrapped = append(unwrapped, elem)
			return nil
		}); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = RetrieveMenusByRestaurantOKApplicationJSON(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s RetrieveMenusByRestaurantOKApplicationJSON) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *RetrieveMenusByRestaurantOKApplicationJSON) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ValidateRestaurantDishesBadRequest as json.
func (s *ValidateRestaurantDishesBadRequest) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes ValidateRestaurantDishesBadRequest from json.
func (s *ValidateRestaurantDishesBadRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ValidateRestaurantDishesBadRequest to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = ValidateRestaurantDishesBadRequest(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ValidateRestaurantDishesBadRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ValidateRestaurantDishesBadRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ValidateRestaurantDishesInternalServerError as json.
func (s *ValidateRestaurantDishesInternalServerError) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes ValidateRestaurantDishesInternalServerError from json.
func (s *ValidateRestaurantDishesInternalServerError) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ValidateRestaurantDishesInternalServerError to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = ValidateRestaurantDishesInternalServerError(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ValidateRestaurantDishesInternalServerError) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ValidateRestaurantDishesInternalServerError) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ValidateRestaurantDishesNotFound as json.
func (s *ValidateRestaurantDishesNotFound) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes ValidateRestaurantDishesNotFound from json.
func (s *ValidateRestaurantDishesNotFound) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ValidateRestaurantDishesNotFound to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = ValidateRestaurantDishesNotFound(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ValidateRestaurantDishesNotFound) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ValidateRestaurantDishesNotFound) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ValidationList) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ValidationList) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("restaurant")
		json.EncodeUUID(e, s.Restaurant)
	}
	{
		e.FieldStart("dishes")
		e.ArrStart()
		for _, elem := range s.Dishes {
			json.EncodeUUID(e, elem)
		}
		e.ArrEnd()
	}
}

var jsonFieldsNameOfValidationList = [2]string{
	0: "restaurant",
	1: "dishes",
}

// Decode decodes ValidationList from json.
func (s *ValidationList) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ValidationList to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "restaurant":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := json.DecodeUUID(d)
				s.Restaurant = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"restaurant\"")
			}
		case "dishes":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				s.Dishes = make([]uuid.UUID, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem uuid.UUID
					v, err := json.DecodeUUID(d)
					elem = v
					if err != nil {
						return err
					}
					s.Dishes = append(s.Dishes, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"dishes\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ValidationList")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfValidationList) {
					name = jsonFieldsNameOfValidationList[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ValidationList) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ValidationList) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
