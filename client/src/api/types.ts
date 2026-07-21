export interface EventType {
  id: number
  name: string
  description?: string
  duration: number
}

export interface AvailableSlot {
  startTime: string
  endTime: string
}

export interface AvailabilityResponse {
  eventTypeId: number
  duration: number
  slots: AvailableSlot[]
}

export interface Booking {
  id: number
  eventTypeId: number
  eventType: EventType
  startTime: string
  endTime: string
  guestName: string
  guestEmail: string
  notes?: string
  createdAt: string
}

export interface CreateBookingRequest {
  eventTypeId: number
  startTime: string
  guestName: string
  guestEmail: string
  notes?: string
}
