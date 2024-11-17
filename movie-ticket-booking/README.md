Requirements

1. Multiple cinemas in a city and cinema has multiple halls.
2. Each movie in the cinema can have multiple shows, but one hall will show only one show at a time.
3. Cinema displays all available showtimes of a movie.
4. Users can search movies based on the following four criteria: title, language, genre, release date.
5. Users can make a booking at any cinema hall at the available showtime.
6. Each seat type has a fixed cost. There are three types of seats: silver, gold and platinum.
7. Users can select multiple available seats for a show from a given seating arrangement.
8. There can only be one ticket allocated per seat.
9. No two customers should be able to reserver the same seat.

Main Actors

Admin - Add, remove or update a show and movie
Customer - Book, modify or cancel booking, pay for the tickets they booked
System - Sending notifications for new movies, bookings and cancellations


Class diagram
Seat
    ID string
    Type SeatType
    Cost float64
    IsReserved() bool

SilverSeat
GoldSeat
PlatinumSeat

Movie
    Title string
    Language string
    Genre string
    ReleaseDate time.DateTime
    shows []Show

Show
    StartTime time.Time
    EndTime time.Time
    seats []Seat

Hall
    shows []Show

Cinema
    name string
    city string
    id string
    halls []Hall

Admin
    AddCinema(cinema_id, name, city)
    AddHall(cinema_id, hall_id, name)
    AddMovie(hall_id, movie_id, title, language, genre, release_date)
    AddShow(show_id, start_time, end_time, hall_id, movie_id)
    UpdateShow(show_id, start_time, end_time, hall_id)
    RemoveShow(show_id)

Customer
    SearchMovie() []Movie
    BookTicket()
    UpdateTicket()
    CancelTicket()

System
    SendNotification()

Design Patterns to Use
    Singleton for ensuring a single instance of the movie database.
    Factory Method for creating different types of movies or shows.
    Strategy for pricing strategies.
    Observer for notifications (e.g., sending a notification after a successful booking).
    Decorator for adding additional features like premium seating.