// Copyright 2021 Comcast Cable Communications Management, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validation_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/sebdah/goldie/v2"
	"github.com/xmidt-org/ears/pkg/validation"
)

func TestErrorMessage(t *testing.T) {
	testCases := []struct {
		name string
		err  error
	}{
		{name: "Error", err: &validation.Error{}},

		{
			name: "Error_Err",
			err:  &validation.Error{Err: fmt.Errorf("wrapped error")},
		},
		{name: "Errors", err: &validation.Errors{}},

		{
			name: "Errors_Err",
			err: &validation.Errors{Errs: []error{
				fmt.Errorf("wrapped 1 error"),
				fmt.Errorf("wrapped 2 error"),
				fmt.Errorf("wrapped 3 error"),
			}},
		},

		{name: "ProcessingError", err: &validation.ProcessingError{}},

		{
			name: "ProcessingError_Err",
			err:  &validation.ProcessingError{Err: fmt.Errorf("wrapped error")},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t, goldie.WithTestNameForDir(true))
			g.Assert(t, tc.name, []byte(fmt.Sprint(tc.err)))
			g.Assert(t, tc.name+"_unwrapped", []byte(fmt.Sprint(errors.Unwrap(tc.err))))

		})
	}

}
