import { useState } from 'react'
import { Stack, Title, Text } from '@mantine/core'
import type { EventType, AvailableSlot } from '../api/types'
import { EventTypeList } from './EventTypeList'
import { SlotPicker } from './SlotPicker'
import { BookingForm } from './BookingForm'

type Step = 'list' | 'slots' | 'form'

export function BookingPage() {
  const [step, setStep] = useState<Step>('list')
  const [eventType, setEventType] = useState<EventType | null>(null)
  const [slot, setSlot] = useState<AvailableSlot | null>(null)

  function handleSelectType(et: EventType) {
    setEventType(et)
    setStep('slots')
  }

  function handleSelectSlot(s: AvailableSlot) {
    setSlot(s)
    setStep('form')
  }

  function handleDone() {
    setStep('list')
    setEventType(null)
    setSlot(null)
  }

  return (
    <Stack gap="lg">
      <div>
        <Title order={1}>Call Booking</Title>
        <Text c="dimmed">Запишитесь на звонок</Text>
      </div>

      {step === 'list' && (
        <EventTypeList onSelect={handleSelectType} />
      )}

      {step === 'slots' && eventType && (
        <SlotPicker
          eventType={eventType}
          onSelect={handleSelectSlot}
          onBack={() => setStep('list')}
        />
      )}

      {step === 'form' && eventType && slot && (
        <BookingForm
          eventType={eventType}
          slot={slot}
          onBack={() => setStep('slots')}
          onDone={handleDone}
        />
      )}
    </Stack>
  )
}
