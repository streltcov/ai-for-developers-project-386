import { useState } from 'react'
import { MantineProvider } from '@mantine/core'
import { Notifications } from '@mantine/notifications'
import { DatesProvider } from '@mantine/dates'
import { Container, Group, Anchor } from '@mantine/core'
import { BookingPage } from './components/BookingPage'
import { OwnerPage } from './components/OwnerPage'

type Page = 'public' | 'owner'

function App() {
  const [page, setPage] = useState<Page>('public')

  return (
    <MantineProvider>
      <DatesProvider settings={{ locale: 'ru', firstDayOfWeek: 1 }}>
        <Notifications />
        <Container size="md" py="xl">
          <Group gap="md" mb="lg">
            <Anchor
              component="button"
              size="sm"
              fw={page === 'public' ? 700 : 400}
              onClick={() => setPage('public')}
            >
              Запись на звонок
            </Anchor>
            <Anchor
              component="button"
              size="sm"
              fw={page === 'owner' ? 700 : 400}
              onClick={() => setPage('owner')}
            >
              Панель владельца
            </Anchor>
          </Group>

          {page === 'public' && <BookingPage />}
          {page === 'owner' && <OwnerPage />}
        </Container>
      </DatesProvider>
    </MantineProvider>
  )
}

export default App
