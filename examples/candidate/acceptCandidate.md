# acceptCandidate

Accpet a candidate

**URL** : `/candidate/acceptCandidate/{_id}`

**Method** : `PATCH`

**Data constraints**

Provide the id of the candidate in the url.

**Example** 

```
/candidate/acceptCandidate/5e9b5ec12608343d5cbf8c77"`

```

## Success Response

**Condition** : If everything is OK and the candidate exists in the database.

**Code** : `200 SUCCESS`

**Content example**

```json
{
    "Candidate accepted": "5e9b5ec12608343d5cbf8c77"
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
    "error": "There was an internal problem while accepting the candidate"
}

```
