import { useState, useEffect } from 'react'
import {
  Stack, Title, Text, Button, Group, Table, Modal,
  TextInput, Textarea, NumberInput, Loader,
} from '@mantine/core'
import { notifications } from '@mantine/notifications'
import {
  ownerFetchEventTypes, ownerCreateEventType,
  ownerUpdateEventType, ownerDeleteEventType,
} from '../api/client'
import type { EventType, CreateEventTypeRequest } from '../api/types'

export function EventTypeManager() {
  const [types, setTypes] = useState<EventType[]>([])
  const [loading, setLoading] = useState(true)
  const [modalOpen, setModalOpen] = useState(false)
  const [editing, setEditing] = useState<EventType | null>(null)
  const [form, setForm] = useState<CreateEventTypeRequest>({ name: '', duration: 30 })
  const [saving, setSaving] = useState(false)

  function load() {
    setLoading(true)
    ownerFetchEventTypes()
      .then(setTypes)
      .finally(() => setLoading(false))
  }

  useEffect(() => { load() }, [])

  function openCreate() {
    setEditing(null)
    setForm({ name: '', description: '', duration: 30 })
    setModalOpen(true)
  }

  function openEdit(et: EventType) {
    setEditing(et)
    setForm({ name: et.name, description: et.description || '', duration: et.duration })
    setModalOpen(true)
  }

  async function handleSave() {
    setSaving(true)
    try {
      if (editing) {
        await ownerUpdateEventType(editing.id, form)
        notifications.show({ title: 'Обновлено', message: 'Тип события обновлён', color: 'green' })
      } else {
        await ownerCreateEventType(form)
        notifications.show({ title: 'Создано', message: 'Тип события создан', color: 'green' })
      }
      setModalOpen(false)
      load()
    } catch (e: any) {
      notifications.show({ title: 'Ошибка', message: e.message, color: 'red' })
    } finally {
      setSaving(false)
    }
  }

  async function handleDelete(id: number) {
    try {
      await ownerDeleteEventType(id)
      notifications.show({ title: 'Удалено', message: 'Тип события удалён', color: 'green' })
      load()
    } catch (e: any) {
      notifications.show({ title: 'Ошибка', message: e.message, color: 'red' })
    }
  }

  if (loading) return <Loader />

  return (
    <Stack gap="md">
      <Group justify="space-between">
        <Title order={3}>Типы событий</Title>
        <Button onClick={openCreate}>+ Создать</Button>
      </Group>

      {types.length === 0 ? (
        <Text c="dimmed">Нет типов событий</Text>
      ) : (
        <Table striped highlightOnHover>
          <Table.Thead>
            <Table.Tr>
              <Table.Th>Название</Table.Th>
              <Table.Th>Описание</Table.Th>
              <Table.Th>Длительность</Table.Th>
              <Table.Th></Table.Th>
            </Table.Tr>
          </Table.Thead>
          <Table.Tbody>
            {types.map((et) => (
              <Table.Tr key={et.id}>
                <Table.Td>{et.name}</Table.Td>
                <Table.Td c="dimmed">{et.description || '—'}</Table.Td>
                <Table.Td>{et.duration} мин</Table.Td>
                <Table.Td>
                  <Group gap="xs">
                    <Button variant="subtle" size="xs" onClick={() => openEdit(et)}>
                      Изм.
                    </Button>
                    <Button variant="subtle" size="xs" color="red" onClick={() => handleDelete(et.id)}>
                      Удал.
                    </Button>
                  </Group>
                </Table.Td>
              </Table.Tr>
            ))}
          </Table.Tbody>
        </Table>
      )}

      <Modal
        opened={modalOpen}
        onClose={() => setModalOpen(false)}
        title={editing ? 'Редактировать' : 'Создать тип события'}
      >
        <Stack gap="md">
          <TextInput
            label="Название"
            value={form.name}
            onChange={(e) => setForm({ ...form, name: e.currentTarget.value })}
            required
          />
          <Textarea
            label="Описание"
            value={form.description || ''}
            onChange={(e) => setForm({ ...form, description: e.currentTarget.value })}
          />
          <NumberInput
            label="Длительность (мин)"
            value={form.duration}
            onChange={(val) => setForm({ ...form, duration: Number(val) || 30 })}
            min={5}
            step={5}
          />
          <Group justify="flex-end">
            <Button variant="subtle" onClick={() => setModalOpen(false)}>Отмена</Button>
            <Button onClick={handleSave} loading={saving}>Сохранить</Button>
          </Group>
        </Stack>
      </Modal>
    </Stack>
  )
}
