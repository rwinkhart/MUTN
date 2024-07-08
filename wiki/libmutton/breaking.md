## Planned Breaking Changes
Leading up to the v1.0.0 release, breaking changes are both expected and planned. These changes are expected to require manual intervention from the end-user, and thus MUTN/libmutton should not be used prior to v1.0.0 if this is not acceptable. 

These changes include, but may expand beyond the following:

- Migration to Go-native encryption (no reliance on GnuPG)
  - Will be based on symmetrical encryption
  - Will eventually allow combining multiple common encryption algorithms (cascading encryption)
- Password aging data will be stored for each entry (to remind the user when it is time to change passwords)
  - Will be included in entry names or in external file (to prevent needing to decrypt entries to access this information)