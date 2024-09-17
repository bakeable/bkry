package portal_event

var ExampleJSON = `{
    entityId: string, // The entity ID
    entityName: string, // The entity name
    _kind: string, // The entity kind
    metadata: Record<string, any>, // The metadata of the event
    timestamp: number, // The timestamp of the event
    timezone: number, // The timezone of the event
    type: PortalEventType, // The type of the event
    userId: string, // The user triggering the event
}`