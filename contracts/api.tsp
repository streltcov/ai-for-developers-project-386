import "@typespec/http";
import "@typespec/rest";

using TypeSpec.Http;
using TypeSpec.Rest;

@service({
  title: "Call Booking",
  version: "1.0.0",
})
@doc("Educational project: simplified time booking service (inspired by Cal.com)")
namespace BookingService;

/**
 * ============================================
 *             DOMAIN ENTITIES
 * ============================================
 */

// Event Type
@doc("Event type that a guest can book")
model EventType {
  @doc("Unique identifier of the event type")
  id: int32;

  @doc("Name of the event type")
  name: string;

  @doc("Description of the event type")
  description?: string;

  @doc("Duration of the event in minutes")
  duration: int32;
}

// Booking
@doc("Booking record for a meeting")
model Booking {
  @doc("Unique identifier of the booking")
  id: int32;

  @doc("Identifier of the event type")
  eventTypeId: int32;

  @doc("Event type (nested object)")
  eventType: EventType;

  @doc("Start time of the meeting (UTC)")
  startTime: utcDateTime;

  @doc("End time of the meeting (UTC)")
  endTime: utcDateTime;

  @doc("Guest's name")
  guestName: string;

  @doc("Guest's email")
  guestEmail: string;

  @doc("Additional notes from the guest")
  notes?: string;

  @doc("Date and time when the booking was created")
  createdAt: utcDateTime;
}

// Available Slot (computed)
@doc("Available time slot (computed dynamically)")
model AvailableSlot {
  @doc("Start time of the slot")
  startTime: utcDateTime;

  @doc("End time of the slot")
  endTime: utcDateTime;
}

// Availability response
@doc("Response containing available slots")
model AvailabilityResponse {
  @doc("Identifier of the event type")
  eventTypeId: int32;

  @doc("Duration of the event in minutes")
  duration: int32;

  @doc("List of available slots")
  slots: AvailableSlot[];
}

/**
 * ============================================
 *             PUBLIC ENDPOINTS (GUEST)
 * ============================================
 */
@route("/api/public")
@tag("Public")
namespace Public {

  @doc("Get list of all event types (for guests)")
  @get
  @route("/event-types")
  op listEventTypes(): EventType[];

  @doc("Get available slots for the next 14 days")
  @get
  @route("/availability")
  op getAvailability(
    @query @doc("Identifier of the event type") eventTypeId: int32,
    @query @doc("Start date of the period (YYYY-MM-DD)") from?: plainDate,
    @query @doc("End date of the period (YYYY-MM-DD)") to?: plainDate,
  ): AvailabilityResponse;

  @doc("Create a booking for the selected slot")
  @post
  @route("/bookings")
  op createBooking(@body request: {
    @doc("Identifier of the event type")
    eventTypeId: int32;

    @doc("Start time (must be available)")
    startTime: utcDateTime;

    @doc("Guest's name")
    guestName: string;

    @doc("Guest's email")
    guestEmail: string;

    @doc("Additional notes")
    notes?: string;
  }): Booking;
}

/**
 * ============================================
 *             OWNER ENDPOINTS
 * ============================================
 */
@route("/api")
@tag("Owner")
namespace Owner {

  // --- Event Types Management ---

  @doc("Get all event types")
  @get
  @route("/event-types")
  op listEventTypes(): EventType[];

  @doc("Create a new event type")
  @post
  @route("/event-types")
  op createEventType(@body body: {
    name: string;
    description?: string;
    duration: int32;
  }): EventType;

  @doc("Get a single event type")
  @get
  @route("/event-types/{id}")
  op getEventType(@path id: int32): EventType;

  @doc("Update an event type")
  @put
  @route("/event-types/{id}")
  op updateEventType(
    @path id: int32,
    @body body: {
      name?: string;
      description?: string;
      duration?: int32;
    }
  ): EventType;

  @doc("Delete an event type")
  @delete
  @route("/event-types/{id}")
  op deleteEventType(@path id: int32): void;

  // --- Bookings ---

  @doc("Get list of all upcoming bookings (across all event types)")
  @get
  @route("/bookings")
  op listBookings(
    @query @doc("Filter from date") from?: plainDate,
    @query @doc("Filter to date") to?: plainDate,
    @query @doc("Filter by event type") eventTypeId?: int32,
  ): Booking[];
}