import { MantineProvider } from '@mantine/core'
import { Notifications } from '@mantine/notifications'
import { DatesProvider } from '@mantine/dates'
import { Container } from '@mantine/core'
import { BookingPage } from './components/BookingPage'

function App() {
  return (
    <MantineProvider>
      <DatesProvider settings={{ locale: 'ru', firstDayOfWeek: 1 }}>
        <Notifications />
        <Container size="md" py="xl">
          <BookingPage />
        </Container>
      </DatesProvider>
    </MantineProvider>
  )
}

export default App
