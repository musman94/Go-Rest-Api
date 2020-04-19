# denyCandidate

Deny a candidate

**URL** : `/candidate/denyCandidate/{_id}`

**Method** : `PATCH`

**Data constraints**

Provide the id of the candidate in the url.

**Example** 

```
 /assignee/denyCandidate/5e9b5ec12608343d5cbf8c77"`

```

## Success Response

**Condition** : If everything is OK and the candidate exists in the database.

**Code** : `200 SUCCESS`

**Content example**

```json
{
    "Candidate denied": "5e9b5ec12608343d5cbf8c77"
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
