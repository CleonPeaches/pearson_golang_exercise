# Programming Assignment: JSON Parsing for Quiz Questions

## Problem Statement

You are tasked with developing a Go program to parse a nested JSON file containing quiz questions and answers. Each question in the JSON file has multiple choice options, with one correct answer and several incorrect distractor options. Your program should allow a student to select a specific topic from the list of provided topics and retrieve questions from the JSON file based on that topic. The program should accept user answers and save them to the database, responding with whether the answer was correct or incorrect.

## Requirements

### JSON File Structure

The JSON file contains an array of objects, each representing a quiz question.

Each question object has the following structure:

[
{
"question": "What is the capital of France?",
"options": ["Paris", "London", "Berlin", "Rome"],
"correct_answer": "Paris",
"distractors": ["London", "Berlin", "Rome"]
},
{
"question": "Who wrote 'Romeo and Juliet'?",
"options": ["William Shakespeare", "Jane Austen", "Charles Dickens", "Leo Tolstoy"],
"correct_answer": "William Shakespeare",
"distractors": ["Jane Austen", "Charles Dickens", "Leo Tolstoy"]
}
]

### Example

If the user selects keywords for an interested topic, e.g., "France", the program should display:

**Question:** What is the capital of France?

**Options:**

1. Paris
2. London
3. Berlin
4. Rome

## Program Functionality

1. The program should read the JSON file and parse its contents into a suitable data structure and save it. DONE
2. The program should accept user input for the selected topic. DONE
3. The program should search the questions for any question containing the keyword/topic in the question text. DONE
4. If a matching question is found, the program should display the questions along with its options to the user. DONE
5. If no matching question is found, the program should display a message indicating that no questions were found for the given keyword.
6. The program should accept user answers and respond back with correct or incorrect.
7. (Added per problem statement) The program should save user answers to the mock database.

## Input/Output

- The API should have valid headers.
- Input validation should be performed to ensure that the keyword entered by the user is not empty.

## Error Handling

The program should handle any potential errors gracefully, such as file not found, invalid JSON format, incorrect user input, network errors, etc.

## Unit Tests

- The API should have unit tests to test the code.

## Database

- What type of database would you pick (relational, key-value, document-based, graph-based)?
- Indexing strategy.
- Partitioning and scaling strategy.
