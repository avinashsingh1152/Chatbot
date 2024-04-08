import google.generativeai as palm

palm.configure(api_key="AIzaSyBh2vOXx9sOrRVVGF0pSq5OACBA87V7cps")


def generateText(userInput):
    models = [m for m in palm.list_models() if 'generateText' in m.supported_generation_methods]
    model = models[0].name

    inputPrompt = "can you please response as a chat for for question asked " + userInput
    completions = palm.generate_text(model=model, prompt=inputPrompt, temperature=0.7, max_output_tokens=200)

    ans = completions.result
    return ans
