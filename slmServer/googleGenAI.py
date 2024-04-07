import google.generativeai as palm

palm.configure(api_key = "AIzaSyBh2vOXx9sOrRVVGF0pSq5OACBA87V7cps")

models = [m for m in palm.list_models() if 'generateText' in m.supported_generation_methods]

model = models[0].name

userInput = input("enter: ")
completions = palm.generate_text(model = model, prompt = userInput, temperature = 0.5, max_output_tokens = 200);

ans = completions.result
print(ans)