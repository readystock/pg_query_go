// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/* ----------------
 *	 BitmapOr node -
 *		Generate the union of the results of sub-plans.
 *
 * The subplans must be of types that yield tuple bitmaps.  The targetlist
 * and qual fields of the plan are unused and are always NIL.
 * ----------------
 */
type BitmapOr struct {
	Plan        Plan   `json:"plan"`
	Bitmapplans []Node `json:"bitmapplans"`
}

func (node BitmapOr) MarshalJSON() ([]byte, error) {
	type BitmapOrMarshalAlias BitmapOr
	return json.Marshal(map[string]interface{}{
		"BITMAPOR": (*BitmapOrMarshalAlias)(&node),
	})
}

func (node *BitmapOr) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["plan"] != nil {
		err = json.Unmarshal(fields["plan"], &node.Plan)
		if err != nil {
			return
		}
	}

	if fields["bitmapplans"] != nil {
		node.Bitmapplans, err = UnmarshalNodeArrayJSON(fields["bitmapplans"])
		if err != nil {
			return
		}
	}

	return
}
