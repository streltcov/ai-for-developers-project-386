import { MantineProvider } from '@mantine/core'
import { Notifications } from '@mantine/notifications'
import { DatesProvider } from '@mantine/dates'
import { Text, Title, Container, Stack } from '@mantine/core'

function App() {
  return (
    <MantineProvider>
      <DatesProvider settings={{ locale: 'ru', firstDayOfWeek: 1 }}>
        <Notifications />
        <Container size="md" py="xl">
          <Stack gap="lg">
            <Title order={1}>Call Booking</Title>
            <Text c="dimmed">Сервис бронирования звонков</Text>
          </Stack>
        </Container>
      </DatesProvider>
    </MantineProvider>
  )
}

export default App
