import { useState, useEffect } from 'react'
import { SimpleGrid, Text, Loader, Stack, Title } from '@mantine/core'
import { fetchEventTypes } from '../api/client'
import type { EventType } from '../api/types'
import { EventTypeCard } from './EventTypeCard'

interface Props {
  onSelect: (et: EventType) => void
}

export function EventTypeList({ onSelect }: Props) {
  const [types, setTypes] = useState<EventType[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    fetchEventTypes()
      .then(setTypes)
      .catch((e) => setError(e.message))
      .finally(() => setLoading(false))
  }, [])

  if (loading) return <Loader />
  if (error) return <Text c="red">{error}</Text>
  if (!types || types.length === 0) return <Text c="dimmed">Нет доступных типов событий</Text>

  return (
    <Stack gap="md">
      <Title order={2}>Выберите тип события</Title>
      <SimpleGrid cols={{ base: 1, sm: 2 }}>
        {types.map((et) => (
          <EventTypeCard key={et.id} eventType={et} onSelect={onSelect} />
        ))}
      </SimpleGrid>
    </Stack>
  )
}
