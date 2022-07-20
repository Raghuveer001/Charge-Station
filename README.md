# Charge-Station
Charging Stations

I have created a Charging Station Management application with below functionalities.

- Add Charging Station

- Start Charging


- Import the necessary packages

   - "encoding/json "           - Go core package for handling JSON data
   - "net/http"                 - a Go HTTP package for handling HTTP requesting when creating GO APIs
   - "github.com/gorilla/mux"   - for URL matcher and routing. It helps in implementing request routers and match each incoming request with its matching handlers

 
- First run go mod init. This will generate a go mod file that saves the third-party libraries we require.
- To install mux run go get github.com/mux
- Add structs for Add Charging and Start Charging
- Initialize the mux router, then add the router handlers to establish the endpoints of our API.
- Add a port to serve the application

- Add Charging Station
 POST-For adding the Charging Stations

Route - 127.0.0.1:8000/chst

{"StationID":1,"EnergyOutput":"50kwh","Type":"DC"}
{"StationID":2,"EnergyOutput":"10kwh","Type":"AC"}
{"StationID":3,"EnergyOutput":"100kwh","Type":"AC"}
{"StationID":4,"EnergyOutput":"150kwh","Type":"DC"}

GET - To view all the added Charging Stations

Route- 127.0.0.1:8000/allchst

[
    {
        "StationID": 1,
        "EnergyOutput": "50kwh",
        "Type": "DC"
    },
    {
        "StationID": 2,
        "EnergyOutput": "10kwh",
        "Type": "AC"
    },
    {
        "StationID": 3,
        "EnergyOutput": "100kwh",
        "Type": "AC"
    },
    {
        "StationID": 4,
        "EnergyOutput": "150kwh",
        "Type": "DC"
    }
]


POST - To Add Vehicle battery capacity, vehicle charge and Charging start time

Route - 127.0.0.1:8000/stch

 {"StationID":1,"VehicleBatteryCapacity":"500kwh","CurrentVehicleCharge":"100KWH","ChargingStartTime":"2022-07-16T13:30:002"}
 {"StationID":2,"VehicleBatteryCapacity":"450kwh","CurrentVehicleCharge":"200KWH","ChargingStartTime":"2022-07-16T14:30:002"}
 {"StationID":3,"VehicleBatteryCapacity":"500kwh","CurrentVehicleCharge":"100KWH","ChargingStartTime":"2022-07-16T14:30:002"}
 {"StationID":4,"VehicleBatteryCapacity":"500kwh","CurrentVehicleCharge":"100KWH","ChargingStartTime":"2022-07-16T14:30:002"}

GET - For Start Charging

Route - 127.0.0.1:8000/allstch
[
    {
        "StationID": 1,
        "VehicleBatteryCapacity": "500kwh",
        "CurrentVehicleCharge": "100KWH",
        "ChargingStartTime": "2022-07-16T13:30:002"
    },
    {
        "StationID": 2,
        "VehicleBatteryCapacity": "450kwh",
        "CurrentVehicleCharge": "200KWH",
        "ChargingStartTime": "2022-07-16T14:30:002"
    },
    {
        "StationID": 3,
        "VehicleBatteryCapacity": "500kwh",
        "CurrentVehicleCharge": "100KWH",
        "ChargingStartTime": "2022-07-16T14:30:002"
    },
    {
        "StationID": 4,
        "VehicleBatteryCapacity": "500kwh",
        "CurrentVehicleCharge": "100KWH",
        "ChargingStartTime": "2022-07-16T14:30:002"
    }
]
