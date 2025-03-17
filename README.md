- Travlers
  - type: Adult, child, elderly
  - gender:
  - name:
  - age:
- SearchTravel:
  - travelType: oneway/roundTrip
  - origin: IATA code of airport/country
  - destination:  IATA code of airport/country
  - departureDate:
  - returnDate: only for round trips
- ConfirmTravel:
  - rules: a child can not travel alone
  - travlers: []
  - travelId/Code:
  - travelType:
  - returnDate:
- for roundTrips, do we create another flight?

flights:
code:
travlers:
chekedIn:
confirmed:


// app structure

internal:
  - have all the codes related to our internal business logic implementation
  - handlers:
    - are the gin handlers called from api client, they accept a request, do validation on data, call a domain usecase (main business logic), then return the response to the caller
  - domain:
    - this is where all business layer logics happen, including to calls to pkg methods
  - dto: holds all data transfer objects used
pkg:
  - have implementation of things we rely on such as API integration
  - forexample: the bargain api method which is made to implement the FlightSearcher interface
    - this lets us swith between providers easly on app instantiation