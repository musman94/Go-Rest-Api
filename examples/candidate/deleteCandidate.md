# deleteCandidate

Delete a candidate from the database.

**URL** : `/candidate/deleteCandidate/{_id}`

**Method** : `DELETE`

**Data constraints**

Provide the id of the candidate in the url.

**Example** 

```
/candidate/deleteCandidate/5e9b564ab0d66741ab30fed5"`
```

## Success Response

**Condition** : If everything is OK and the candidate exists in the database.

**Code** : `200 SUCCESS`

**Content example**

```json
{
    "Candidate deleted": "5e9b564ab0d66741ab30fed5"
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

### Or

**Condition** : If there was a problem on the database side

**Code** : `500 BAD REQUEST`

**Content example**

```json
{
    "error": "There was an internal problem while deleting the candidate"
}

```
