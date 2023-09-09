import { describe, it, expect } from 'vitest';

import { mount } from '@vue/test-utils';
import AppTitle from '../AppTitle.vue';

describe('AppTitle', () => {
  it('renders properly', () => {
    const wrapper = mount(AppTitle);
    expect(wrapper.text()).toContain('Trequer');
  });
});
