from transformers import pipeline

#AIzaSyBh2vOXx9sOrRVVGF0pSq5OACBA87V7cps

generator = pipeline("text-generation", model="mistralai/Mistral-7B-Instruct-v0.2")

userInput = input("Enter a sentence: ")
output = generator(userInput, temperature=0.8)

print(output)