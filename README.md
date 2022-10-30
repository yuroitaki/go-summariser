# go-summariser
A http server that uses GPT-3 API to summarise a body of text.

## Features
- Summarise a body of text using GPT-3 API with custom temperature, engine setting
- Store summarised result in a database to allow caching

## Environment File Format
```yaml
OPENAI_API_KEY=key
POSTGRES_DB=db
POSTGRES_PASSWORD=pwd
POSTGRES_USER=user
```

## Examples
Curl command:
```bash
curl --location --request POST 'http://localhost:55100/summariser' \
--header 'Content-Type: application/json' \
```
Request body:
```json
{
    "Text": "An old man lived in the village. The whole village was tired of him; he was always gloomy, he constantly complained and was always in a bad mood. The longer he lived, the viler he became and more poisonous were his words. People did their best to avoid him because his misfortune was contagious. He created the feeling of unhappiness in others. But one day, when he turned eighty, an incredible thing happened. Instantly everyone started hearing the rumor: 'The old man is happy today, he doesn't complain about anything, smiles, and even his face is freshened up.' The whole village gathered around the man and asked him, 'What happened to you?' The old man replied, 'Nothing special. Eighty years I've been chasing happiness and it was useless. And then I decided to live without happiness and just enjoy life. That's why I'm happy now.'",
    "Temperature": 1,
    "Engine": "text-davinci-001",
    "TopP": 1
}
```
Response body:
```json
{
    "Summary": "An old man who always complained and was always in a bad mood became happy when he turned eighty and decided to just enjoy life.",
    "Text": "An old man lived in the village. The whole village was tired of him; he was always gloomy, he constantly complained and was always in a bad mood. The longer he lived, the viler he became and more poisonous were his words. People did their best to avoid him because his misfortune was contagious. He created the feeling of unhappiness in others. But one day, when he turned eighty, an incredible thing happened. Instantly everyone started hearing the rumor: 'The old man is happy today, he doesn't complain about anything, smiles, and even his face is freshened up.' The whole village gathered around the man and asked him, 'What happened to you?' The old man replied, 'Nothing special. Eighty years I've been chasing happiness and it was useless. And then I decided to live without happiness and just enjoy life. That's why I'm happy now.'",
    "Temperature": 1,
    "Engine": "text-davinci-001",
    "TopP": 1
}
```