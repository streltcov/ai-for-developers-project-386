import { useState, useEffect } from 'react'
import {
  Stack, Title, Text, Button, Group, Loader, SimpleGrid,
  UnstyledButton,
} from '@mantine/core'
import dayjs from 'dayjs'
import utc from 'dayjs/plugin/utc'
import { fetchAvailability } from '../api/client'
import type { EventType, AvailableSlot } from '../api/types'

dayjs.extend(utc)

interface Props {
  eventType: EventType
  onSelect: (slot: AvailableSlot) => void
  onBack: () => void
}

export function SlotPicker({ eventType, onSelect, onBack }: Props) {
  const [slots, setSlots] = useState<AvailableSlot[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)
  const [selectedDate, setSelectedDate] = useState<string>(
    dayjs().utc().format('YYYY-MM-DD')
  )

  useEffect(() => {
    setLoading(true)
    fetchAvailability(eventType.id)
      .then((res) => setSlots(res.slots))
      .catch((e) => setError(e.message))
      .finally(() => setLoading(false))
  }, [eventType.id])

  const dates = [...new Set(slots.map((s) => dayjs(s.startTime).format('YYYY-MM-DD')))]
  const slotsForDate = slots.filter(
    (s) => dayjs(s.startTime).format('YYYY-MM-DD') === selectedDate
  )

  return (
    <Stack gap="md">
      <Group>
        <Button variant="subtle" onClick={onBack}>← Назад</Button>
        <Title order={2}>{eventType.name}</Title>
      </Group>
      <Text c="dimmed">Длительность: {eventType.duration} мин</Text>

      {loading && <Loader />}
      {error && <Text c="red">{error}</Text>}

      {!loading && !error && (
        <>
          <Title order={4}>Выберите дату</Title>
          <Group>
            {dates.map((d) => (
              <UnstyledButton
                key={d}
                p="sm"
                style={{
                  borderRadius: 8,
                  border: d === selectedDate
                    ? '2px solid var(--mantine-color-blue-6)'
                    : '1px solid var(--mantine-color-gray-3)',
                  background: d === selectedDate
                    ? 'var(--mantine-color-blue-0)'
                    : undefined,
                }}
                onClick={() => setSelectedDate(d)}
              >
                <Stack gap={2} align="center">
                  <Text size="xs" c="dimmed">{dayjs(d).format('dd')}</Text>
                  <Text fw={600}>{dayjs(d).format('D')}</Text>
                  <Text size="xs" c="dimmed">{dayjs(d).format('MMM')}</Text>
                </Stack>
              </UnstyledButton>
            ))}
          </Group>

          {dates.length === 0 && (
            <Text c="dimmed">Нет свободных слотов на ближайшие 14 дней</Text>
          )}

          {slotsForDate.length > 0 && (
            <>
              <Title order={4}>Время</Title>
              <SimpleGrid cols={{ base: 3, sm: 5, md: 7 }}>
                {slotsForDate.map((slot) => (
                  <Button
                    key={slot.startTime}
                    variant="light"
                    size="sm"
                    onClick={() => onSelect(slot)}
                  >
                    {dayjs(slot.startTime).format('HH:mm')}
                  </Button>
                ))}
              </SimpleGrid>
            </>
          )}
        </>
      )}
    </Stack>
  )
}
