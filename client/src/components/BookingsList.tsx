import { useState, useEffect } from 'react'
import {
  Stack, Title, Text, Table, Loader, Badge,
} from '@mantine/core'
import dayjs from 'dayjs'
import utc from 'dayjs/plugin/utc'
import { ownerFetchBookings } from '../api/client'
import type { Booking } from '../api/types'

dayjs.extend(utc)

export function BookingsList() {
  const [bookings, setBookings] = useState<Booking[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    ownerFetchBookings()
      .then(setBookings)
      .finally(() => setLoading(false))
  }, [])

  if (loading) return <Loader />

  return (
    <Stack gap="md">
      <Title order={3}>Предстоящие встречи</Title>

      {(!bookings || bookings.length === 0) ? (
        <Text c="dimmed">Нет предстоящих бронирований</Text>
      ) : (
        <Table striped highlightOnHover>
          <Table.Thead>
            <Table.Tr>
              <Table.Th>Тип</Table.Th>
              <Table.Th>Дата и время</Table.Th>
              <Table.Th>Гость</Table.Th>
              <Table.Th>Email</Table.Th>
            </Table.Tr>
          </Table.Thead>
          <Table.Tbody>
            {bookings.map((b) => (
              <Table.Tr key={b.id}>
                <Table.Td>
                  <Badge variant="light">{b.eventType.name}</Badge>
                </Table.Td>
                <Table.Td>
                  {dayjs(b.startTime).format('D MMM YYYY, HH:mm')} —{' '}
                  {dayjs(b.endTime).format('HH:mm')}
                </Table.Td>
                <Table.Td>{b.guestName}</Table.Td>
                <Table.Td c="dimmed">{b.guestEmail}</Table.Td>
              </Table.Tr>
            ))}
          </Table.Tbody>
        </Table>
      )}
    </Stack>
  )
}
