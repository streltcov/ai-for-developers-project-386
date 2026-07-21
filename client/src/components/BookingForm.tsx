import { useState } from 'react'
import {
  Stack, Title, Text, Button, Group, TextInput, Textarea,
} from '@mantine/core'
import { notifications } from '@mantine/notifications'
import dayjs from 'dayjs'
import utc from 'dayjs/plugin/utc'
import { createBooking } from '../api/client'
import type { EventType, AvailableSlot } from '../api/types'

dayjs.extend(utc)

interface Props {
  eventType: EventType
  slot: AvailableSlot
  onBack: () => void
  onDone: () => void
}

export function BookingForm({ eventType, slot, onBack, onDone }: Props) {
  const [name, setName] = useState('')
  const [email, setEmail] = useState('')
  const [notes, setNotes] = useState('')
  const [loading, setLoading] = useState(false)

  const start = dayjs(slot.startTime)
  const end = dayjs(slot.endTime)

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault()
    setLoading(true)
    try {
      await createBooking({
        eventTypeId: eventType.id,
        startTime: slot.startTime,
        guestName: name,
        guestEmail: email,
        notes: notes || undefined,
      })
      notifications.show({
        title: 'Бронирование создано',
        message: `${start.format('D MMMM YYYY, HH:mm')} — ${end.format('HH:mm')}`,
        color: 'green',
      })
      onDone()
    } catch (err: any) {
      notifications.show({
        title: 'Ошибка',
        message: err.message,
        color: 'red',
      })
    } finally {
      setLoading(false)
    }
  }

  return (
    <Stack gap="md" component="form" onSubmit={handleSubmit}>
      <Group>
        <Button variant="subtle" onClick={onBack} type="button">← Назад</Button>
        <Title order={2}>Подтвердите запись</Title>
      </Group>

      <Text>
        <strong>{eventType.name}</strong> — {start.format('D MMMM YYYY')}
      </Text>
      <Text c="dimmed">
        {start.format('HH:mm')} — {end.format('HH:mm')} ({eventType.duration} мин)
      </Text>

      <TextInput
        label="Ваше имя"
        placeholder="Иван Иванов"
        value={name}
        onChange={(e) => setName(e.currentTarget.value)}
        required
      />

      <TextInput
        label="Email"
        placeholder="ivan@example.com"
        type="email"
        value={email}
        onChange={(e) => setEmail(e.currentTarget.value)}
        required
      />

      <Textarea
        label="Комментарий (необязательно)"
        placeholder="Тема звонка, вопросы..."
        value={notes}
        onChange={(e) => setNotes(e.currentTarget.value)}
      />

      <Group justify="flex-end">
        <Button type="submit" loading={loading}>
          Записаться
        </Button>
      </Group>
    </Stack>
  )
}
