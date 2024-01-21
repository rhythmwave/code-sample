<?php
/*
Plugin Name: Last Modified Posts Notice
Description: Display a dismissible admin notice showcasing the titles of the last 10 modified and published blog posts.
Version: 1.0
Author: Bayu Dian Nugroho
*/

// Function to display admin notice
function display_last_modified_posts_notice() {
    // Get the last 10 modified and published posts
    $args = array(
        'post_type'      => 'post',
        'post_status'    => 'publish',
        'orderby'        => 'modified',
        'order'          => 'DESC',
        'posts_per_page' => 10,
    );

    $query = new WP_Query($args);

    // Check if there are posts
    if ($query->have_posts()) {
        $msg = 'Last 10 Modified and Published Posts:<br>';

        // Loop through the posts and append titles to the message
        while ($query->have_posts()) {
            $query->the_post();
            $msg .= get_the_title() . '<br>';
        }

        // Display the admin notice
        ?>
        <div class="notice notice-success is-dismissible">
            <p><?php echo $msg; ?></p>
        </div>
        <?php
    }

    // Reset post data
    wp_reset_postdata();
}

// Hook to display the notice in the admin area
add_action('admin_notices', 'display_last_modified_posts_notice');
