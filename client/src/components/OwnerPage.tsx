import { useState } from 'react'
import { Stack, Title, Text, Tabs } from '@mantine/core'
import { EventTypeManager } from './EventTypeManager'
import { BookingsList } from './BookingsList'

export function OwnerPage() {
  const [tab, setTab] = useState<string | null>('events')

  return (
    <Stack gap="lg">
      <div>
        <Title order={1}>Панель владельца</Title>
        <Text c="dimmed">Управление событиями и бронированиями</Text>
      </div>

      <Tabs value={tab} onChange={setTab}>
        <Tabs.List>
          <Tabs.Tab value="events">Типы событий</Tabs.Tab>
          <Tabs.Tab value="bookings">Встречи</Tabs.Tab>
        </Tabs.List>

        <Tabs.Panel value="events" pt="md">
          <EventTypeManager />
        </Tabs.Panel>

        <Tabs.Panel value="bookings" pt="md">
          <BookingsList />
        </Tabs.Panel>
      </Tabs>
    </Stack>
  )
}
