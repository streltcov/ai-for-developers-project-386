import type {
  EventType, AvailabilityResponse, Booking, CreateBookingRequest,
  CreateEventTypeRequest, UpdateEventTypeRequest,
} from './types'

async function request<T>(url: string, options?: RequestInit): Promise<T> {
  const res = await fetch(url, options)
  if (!res.ok) {
    const body = await res.json().catch(() => ({ error: res.statusText }))
    throw new Error(body.error || res.statusText)
  }
  if (res.status === 204) return undefined as T
  return res.json()
}

// --- Public ---

const PUBLIC = '/api/public'

export function fetchEventTypes(): Promise<EventType[]> {
  return request<EventType[]>(`${PUBLIC}/event-types`)
}

export function fetchAvailability(eventTypeId: number, from?: string, to?: string): Promise<AvailabilityResponse> {
  const params = new URLSearchParams({ eventTypeId: String(eventTypeId) })
  if (from) params.set('from', from)
  if (to) params.set('to', to)
  return request<AvailabilityResponse>(`${PUBLIC}/availability?${params}`)
}

export function createBooking(data: CreateBookingRequest): Promise<Booking> {
  return request<Booking>(`${PUBLIC}/bookings`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(data),
  })
}

// --- Owner ---

const OWNER = '/api'

export function ownerFetchEventTypes(): Promise<EventType[]> {
  return request<EventType[]>(`${OWNER}/event-types`)
}

export function ownerCreateEventType(data: CreateEventTypeRequest): Promise<EventType> {
  return request<EventType>(`${OWNER}/event-types`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(data),
  })
}

export function ownerUpdateEventType(id: number, data: UpdateEventTypeRequest): Promise<EventType> {
  return request<EventType>(`${OWNER}/event-types/${id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(data),
  })
}

export function ownerDeleteEventType(id: number): Promise<void> {
  return request<void>(`${OWNER}/event-types/${id}`, {
    method: 'DELETE',
  })
}

export function ownerFetchBookings(from?: string, to?: string, eventTypeId?: number): Promise<Booking[]> {
  const params = new URLSearchParams()
  if (from) params.set('from', from)
  if (to) params.set('to', to)
  if (eventTypeId) params.set('eventTypeId', String(eventTypeId))
  const qs = params.toString()
  return request<Booking[]>(`${OWNER}/bookings${qs ? `?${qs}` : ''}`)
}
