# readCandidate

Fetch a candidate from the database.

**URL** : `/candidate/readCandidate/{_id}`

**Method** : `GET`

**Data constraints**

Provide the id of the candidate in the url.

**Example** 

```
 /assignee/readCandidate/5e9b564ab0d66741ab30fed5"`
```

## Success Response

**Condition** : If everything is OK and the candidate exists in the database.

**Code** : `200 SUCCESS`

**Content example**

```json
{
    "_id": "5e9b564ab0d66741ab30fed5",
    "first_name": "mehmet",
    "last_name": "halit",
    "email": "mehmethalit@gmail.com",
    "department": "Development",
    "university": "Bilkent University",
    "experience": true,
    "status": "Pending",
    "meeting_count": 0,
    "next_meeting": null,
    "assignee": "5e9b2e37a358a83e95401984",
    "application_date": "18-04-2020"
}
```

## Error Response

**Condition** : If the candidate doesn't exist in the database.

**Code** : `404 NOT FOUND`

**Content example** :

```json
{
    "error": "Invalid Candidate ID"
}
```
