import type { EventType, AvailabilityResponse, Booking, CreateBookingRequest } from './types'

const BASE = '/api/public'

async function request<T>(url: string, options?: RequestInit): Promise<T> {
  const res = await fetch(url, options)
  if (!res.ok) {
    const body = await res.json().catch(() => ({ error: res.statusText }))
    throw new Error(body.error || res.statusText)
  }
  if (res.status === 204) return undefined as T
  return res.json()
}

export function fetchEventTypes(): Promise<EventType[]> {
  return request<EventType[]>(`${BASE}/event-types`)
}

export function fetchAvailability(eventTypeId: number, from?: string, to?: string): Promise<AvailabilityResponse> {
  const params = new URLSearchParams({ eventTypeId: String(eventTypeId) })
  if (from) params.set('from', from)
  if (to) params.set('to', to)
  return request<AvailabilityResponse>(`${BASE}/availability?${params}`)
}

export function createBooking(data: CreateBookingRequest): Promise<Booking> {
  return request<Booking>(`${BASE}/bookings`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(data),
  })
}
