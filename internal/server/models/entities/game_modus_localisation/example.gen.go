package game_modus_localisation

var ExampleJSON = `{
    explanation: string, // Detailed explanation of the game mode
    game_modus_id: string, // The ID of the parent entity for which this is the child document
    language_id: string, // The default language code for the game mode
    sections: {
      key: "", // A fixed key across languages 
      order: "", // The ordering of the explanation sections 
      subtitle: "", // The subtitle of the section 
      text: "", // The text content of the section 
      title: "", // The title of the section 
    }[], // Sections used to elaborate the game mode further
    short: string, // A short summarized explanation of the game mode
    subtitle: string, // The subtitle of the game mode
    title: string, // The title of the game mode
}`