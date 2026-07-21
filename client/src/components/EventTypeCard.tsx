import { Card, Text, Badge, Stack } from '@mantine/core'
import type { EventType } from '../api/types'

interface Props {
  eventType: EventType
  onSelect: (et: EventType) => void
}

export function EventTypeCard({ eventType, onSelect }: Props) {
  return (
    <Card
      shadow="sm"
      padding="lg"
      radius="md"
      withBorder
      style={{ cursor: 'pointer' }}
      onClick={() => onSelect(eventType)}
    >
      <Stack gap="xs">
        <Text fw={600} size="lg">{eventType.name}</Text>
        {eventType.description && (
          <Text c="dimmed" size="sm">{eventType.description}</Text>
        )}
        <Badge variant="light" color="blue">
          {eventType.duration} мин
        </Badge>
      </Stack>
    </Card>
  )
}
