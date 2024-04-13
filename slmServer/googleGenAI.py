import google.generativeai as palm
import os

apiKey = os.environ.get('google_api_key')
palm.configure(api_key=apiKey)

def generateText(userInput):
    models = [m for m in palm.list_models() if 'generateText' in m.supported_generation_methods]
    model = models[0].name

    completions = palm.generate_text(model=model, prompt=userInput, temperature=0.7, max_output_tokens=200)

    ans = completions.result
    return ans
