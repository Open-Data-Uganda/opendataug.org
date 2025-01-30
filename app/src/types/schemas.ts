import { z } from 'zod';

export const LoginSchema = z.object({
  email: z.string().min(1, 'Email is required').email().trim(),
  password: z.string().min(1, 'Password is required')
});

export const ResetPasswordSchema = z.object({
  email: z.string().min(1, 'Email is required').email().trim()
});

export const SignUpSchema = z.object({
  firstName: z.string().min(1, 'First name is required'),
  lastName: z.string().min(1, 'Last name is required'),
  email: z.string().email('Invalid email address'),
  password: z.string().min(8, 'Password must be at least 8 characters')
});

export const APIKeySchema = z.object({
  name: z.string().min(1, 'API Key name is required').max(10, 'API Key name cannot exceed 10 characters').trim()
});

export const InviteUserSchema = z.object({
  email: z.string().min(1, 'Email is required').email().trim(),
  first_name: z.string().min(1, 'First name is required'),
  other_name: z.string().min(1, 'Last name is required'),
  telephone: z.string().optional(),
  role: z.string().min(1, 'Role is required')
});

export const SetPasswordSchema = z.object({
  password: z.string().min(1, 'Password is required'),
  confirm_password: z.string().min(1, 'Confirm password is required')
});

export const EditProfileSchema = z.object({
  first_name: z.string().optional(),
  other_name: z.string().optional(),
  email: z.string().optional(),
  phone: z.string().optional()
});
