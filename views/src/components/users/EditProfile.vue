<template>
  <div class="register-box">
    <div class="register-box-body">
      <h3 class="login-box-msg">Edit Profile</h3>
      <form v-on:submit.prevent="updateProfile()">
        <div v-if="errors.length > 0">
          <div  class="alert alert-danger">
            <a href="#" class="close" data-dismiss="alert">&times;</a>
            <ul v-for="msg in errors">
                <li>{{ msg }}</li>
            </ul>
          </div>
        </div>
        <div class="form-group has-feedback">
          <input type="text" class="form-control" placeholder="Name" v-model="profile.name" required="required">
          <span class="glyphicon glyphicon-user form-control-feedback"></span>
        </div>

        <div class="form-group has-feedback">
          <input type="text" class="form-control" placeholder="Title" v-model="profile.title">
          <span class="glyphicon glyphicon-sunglasses form-control-feedback"></span>
        </div>
        
        <div class="form-group has-feedback">
          <input type="text" class="form-control" placeholder="Contact" v-model="profile.contact">
          <span class="glyphicon glyphicon-phone form-control-feedback"></span>
        </div>

        <div class="form-group has-feedback">
          <input type="text" class="form-control" id="address" v-model="profile.address" placeholder="Address">
          <span class="glyphicon glyphicon-map form-control-feedback"></span>
        </div>

        <div id="map"></div>
        
        <div class="row">
          <!-- /.col -->
          <div class="col-xs-12">
            <button type="submit" class="btn btn-primary btn-block btn-flat">Submit</button>
          </div>
        </div>
      </form>
    </div>
    <!-- /.form-box -->

    <br/>
    <div class="text-center">
        <router-link to="/profile" class="btn btn-primary btn-flat">
             <span class="glyphicon glyphicon-chevron-left"></span> Back
        </router-link>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapState } from 'vuex';
import authService from './../../services/auth';
import userService from './../../services/user';

/* global window, google, document */
export default {
  name: 'EditProfile',
  data() {
    return {
      errors: [],
      profile: this.$store.getters.currentUser || {
        name: null,
        title: null,
        contact: null,
        address: null,
        latitude: null,
        longitude: null,
      },
      map: {},
    };
  },
  mounted() {
    google.maps.event.addDomListener(window, 'load', this.initAutocomplete());
    this.initAutocomplete();
  },
  methods: {
    updateProfile() {
      this.errors = [];
      userService.update(this, this.profile.ID, this.profile, '/profile');
    },
    initAutocomplete() {
      const map = new google.maps.Map(document.getElementById('map'), {
        center: {
          lat: parseFloat(this.profile.latitude) || 0,
          lng: parseFloat(this.profile.longitude) || 0,
        },
        zoom: 15,
        mapTypeId: google.maps.MapTypeId.ROADMAP,
      });

      // Create the search box and link it to the UI element.
      const input = document.getElementById('address');
      const searchBox = new google.maps.places.SearchBox(input);
      map.controls[google.maps.ControlPosition.TOP_LEFT].push(input);

      // Bias the SearchBox results towards current map's viewport.
      map.addListener('bounds_changed', () => {
        searchBox.setBounds(map.getBounds());
      });

      let markers = [];
      // [START region_getplaces]
      // Listen for the event fired when the user selects a prediction and retrieve
      // more details for that place.
      searchBox.addListener('places_changed', () => {
        const places = searchBox.getPlaces();

        if (places.length === 0) {
          return;
        }

        // Clear out the old markers.
        markers.forEach((marker) => {
          marker.setMap(null);
        });
        markers = [];

        // For each place, get the icon, name and location.
        const bounds = new google.maps.LatLngBounds();
        places.forEach((place) => {
          const icon = {
            url: place.icon,
            size: new google.maps.Size(71, 71),
            origin: new google.maps.Point(0, 0),
            anchor: new google.maps.Point(17, 34),
            scaledSize: new google.maps.Size(25, 25),
          };

          // Create a marker for each place.
          markers.push(new google.maps.Marker({
            map,
            icon,
            title: place.name,
            position: place.geometry.location,
          }));

          this.profile.address = place.formatted_address;
          this.profile.latitude = place.geometry.location.lat().toString();
          this.profile.longitude = place.geometry.location.lng().toString();

          if (place.geometry.viewport) {
            // Only geocodes have viewport.
            bounds.union(place.geometry.viewport);
          } else {
            bounds.extend(place.geometry.location);
          }
        });
        map.fitBounds(bounds);
      });
      // [END region_getplaces]
    },
  },
  created() {
    authService.currentUser(this);
  },
  computed: {
    ...mapState([
      'currentUser',
    ]),
    ...mapGetters([
      'currentUser',
    ]),
  },
};
</script>


<style>
#map {
  width: 100%;
  height: 200px;
  border: 1px solid #ccc;
  margin: 15px 0;  
}
</style>