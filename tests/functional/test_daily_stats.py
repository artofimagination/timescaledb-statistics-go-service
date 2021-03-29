import pytest
import json


# getResponse unwraps the data/error from json response.
# @expected shall be set to None only if
# the response result is just to generate a component for a test
# but not actually returning a test result.
def getResponse(responseText, expected=None):
    response = json.loads(responseText)
    if "error" in response:
        error = response["error"]
        if expected is None or \
                (expected is not None and error != expected["error"]):
            pytest.fail(f"Failed to run test.\nDetails: {error}")
        return None
    return response["data"]


dataColumns = ("data", "expected")
createTestData = [
    (
        # Input data
        {
          "data_to_store": [
              {
                  "created_at": "Sat Sep 23 2017 15:38:22.0123 GMT+0630",
                  "viewer_id": "4a8c1833-3e4d-4144-8aab-3133f2bdc132",
                  "data": {
                      "test1": "value1",
                      "test2": "value2"
                  }
              }
          ]
        },
        # Expected
        {
            "error": "",
            "data": "OK"
        }),
    (
        # Input data
        {
          "data_to_store": [
              {
                  "created_at": "Sat Sep 23 2017 15:38:23.0000 GMT+0630",
                  "viewer_id": "4a8c1833-3e4d-4144-8aab-3133f2bdc132",
                  "data": {
                      "test3": "value1",
                      "test4": "value2"
                  }
              },
              {
                  "created_at": "Sat Sep 23 2017 15:38:23.0123 GMT+0630",
                  "viewer_id": "863bda70-a0aa-45fc-bd0c-dedd81515292",
                  "data": {
                      "test5": "value1",
                      "test6": "value2"
                  }
              }
          ]
        },
        # Expected
        {
            "error": "",
            "data": "OK"
        }),
    (
        # Input data
        {
          "data_to_store": [
              {
                  "created_at": "Sat Sep 23 2017 15:38:23.0000 GMT+0630",
                  "viewer_id": "4a8c1833-3e4d-4144-8aab-3133f2bdc132",
                  "data": {
                      "test3": "value1",
                      "test4": "value2"
                  }
              }
          ]
        },
        # Expected
        {
            "error": "Failed to save data",
            "data": ""
        })
]

ids = ['Add single row', 'Add multiple rows', 'Failure']


@pytest.mark.parametrize(dataColumns, createTestData, ids=ids)
def test_AddData(httpConnection, data, expected):
    pass
